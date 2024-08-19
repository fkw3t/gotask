package usecase

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/fkw3t/gotask/internal/enum"
	"github.com/fkw3t/gotask/internal/model"
	"github.com/mergestat/timediff"
)

type ListTaskUseCase struct {
	taskRepo task.TaskRepo
}

func NewListTaskUseCase(taskRepo task.TaskRepo) *ListTaskUseCase {
	return &ListTaskUseCase{taskRepo}
}

func (u *ListTaskUseCase) List() error {
	tasks, err := u.taskRepo.List()
	if err != nil {
		return err
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "id\tname\tdescription\tdeadline\tcreated")
	for i, task := range tasks {
		if i == 0 || task.Status == enum.StatusDone {
			continue
		}

		fmt.Fprintln(
			writer,
			fmt.Sprintf(
				"%v\t%v\t%v\t%v\t%v",
				task.Id,
				task.Name,
				transformEmptyFields(task.Description),
				getTimeDiff(task.Deadline),
				getTimeDiff(task.CreatedAt),
			),
		)
	}

	return nil
}

func (u *ListTaskUseCase) ListWithDetails() error {
	tasks, err := u.taskRepo.List()
	if err != nil {
		return err
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "id\tname\tdescription\tdeadline\texpired\tstatus\tdue_date\tcreated")
	for i, task := range tasks {
		if i == 0 {
			continue
		}

		fmt.Fprintln(
			writer,
			fmt.Sprintf(
				"%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v",
				task.Id,
				task.Name,
				transformEmptyFields(task.Description),
				getTimeDiff(task.Deadline),
				func() string {
					if task.Deadline != nil && !task.Deadline.IsZero() {
						if task.Deadline.Before(time.Now()) {
							return "true"
						}

						return "false"
					}

					return "-"
				}(),
				task.Status.String(),
				getTimeDiff(task.DueDate),
				getTimeDiff(task.CreatedAt),
			),
		)
	}

	return nil
}

func transformEmptyFields(field string) string {
	if field == "" {
		return "-"
	}

	return field
}

func getTimeDiff(date *time.Time) string {
	if date != nil && !date.IsZero() {
		return timediff.TimeDiff(*date)
	}

	return "-"
}
