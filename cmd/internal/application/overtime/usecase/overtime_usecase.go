package usecase

import (
	"context"

	"github.com/teguh522/payslip/cmd/internal/application/overtime/command"
	"github.com/teguh522/payslip/cmd/internal/application/overtime/dto"
	"github.com/teguh522/payslip/cmd/internal/domain/overtime/entity"
	"github.com/teguh522/payslip/cmd/internal/domain/overtime/repository"
)

type OvertimeUseCase struct {
	overtimeRepo repository.OvertimeRepository
}

func NewOvertimeUseCase(overtimeRepo repository.OvertimeRepository) *OvertimeUseCase {
	return &OvertimeUseCase{
		overtimeRepo: overtimeRepo,
	}
}

func (uc *OvertimeUseCase) Execute(ctx context.Context, cmd command.OvertimeCommand) (*dto.CreateOvertimeResponse, error) {
	overtime, err := entity.NewOvertime(cmd.Date, cmd.Hours, cmd.Description, cmd.CreatedBy, cmd.UpdatedBy, cmd.EmployeeID, cmd.PeriodID)
	if err != nil {
		return nil, err
	}

	err = uc.overtimeRepo.CreateOvertime(ctx, overtime)
	if err != nil {
		return nil, err
	}

	return &dto.CreateOvertimeResponse{
		Status: "overtime created successfully",
	}, nil
}
