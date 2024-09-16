// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"context"

	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/internal/shared"
)

func NewContextWithDeniedValues(ctx context.Context, deniedValues []any) *shared.ContextWithDeniedValues {
	return shared.NewContextWithDeniedValues(ctx, deniedValues)
}
