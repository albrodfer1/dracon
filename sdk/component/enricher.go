package component

import (
	"context"
	"fmt"
)

// RunEnricher runs an enricher after initialising the run context.
func RunEnricher(ctx context.Context, enricher Enricher, opts ...RunnerOption) error {
	return run(
		ctx,
		func(ctx context.Context, cfg *RunnerConfig) error {
			var (
				workflowID = cfg.WorkflowID
				logger     = LoggerFromContext(ctx).With(logKeyComponentType, "enricher")
				store      = cfg.storerConfig.store
			)

			defer func() {
				if err := store.Close(ctx); err != nil {
					logger.With(logKeyError, err.Error()).Error("closing step failed, ignoring...")
				}
			}()

			logger.Debug("preparing to execute enricher component...")
			logger.Debug("preparing to execute read step...")

			findings, err := store.Read(ctx, workflowID)
			if err != nil {
				logger.With(logKeyError, err.Error()).Error("reading step failed")
				return fmt.Errorf("could not read: %w", err)
			}

			logger = logger.With(logKeyNumParsedFindings, len(findings))
			logger.Debug("read step completed!")

			logger.Debug("preparing to execute enricher step...")
			enrichedFindings, err := enricher.Annotate(ctx, findings)
			if err != nil {
				logger.With(logKeyError, err.Error()).Error("enricher step failed")
				return fmt.Errorf("could not enricher: %w", err)
			}

			logger = logger.With(logKeyNumEnrichedFindings, len(enrichedFindings))
			logger.Debug("enricher step completed!")
			logger.Debug("preparing to execute update step...")

			if err := store.Update(ctx, workflowID, enrichedFindings); err != nil {
				logger.With(logKeyError, err.Error()).Error("updating step failed")
				return fmt.Errorf("could not update: %w", err)
			}

			logger.Debug("updated step completed!")
			logger.Debug("enricher component has completed successfully!")

			return nil
		},
		opts...,
	)
}