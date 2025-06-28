package repository

import (
	"context"

	"github.com/teguh522/payslip/cmd/internal/domain/overtime/entity"
)

type OvertimeRepository interface {
	CreateOvertime(ctx context.Context, overtime *entity.Overtime) error
}
