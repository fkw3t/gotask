package filesystem_repo

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	utils "github.com/fkw3t/gotask/internal"
	"github.com/fkw3t/gotask/internal/enum"
	task "github.com/fkw3t/gotask/internal/model"
)

type TaskRepo struct {
	filepath string
	filename string
}

func (r *TaskRepo) init() error {
	_, err := os.Stat(fmt.Sprintf("%v/%v.csv", r.filepath, r.filename))

	if os.IsNotExist(err) {
		file, err := os.Create(fmt.Sprintf("%v/%v.csv", r.filepath, r.filename))
		if err != nil {
			return fmt.Errorf("failed to create storage file: %v\n", err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		err = writer.Write([]string{"id", "name", "description", "status", "deadline", "due_date", "created_at"})
		if err != nil {
			return fmt.Errorf("failed to write headers columns to storage file: %v\n", err)
		}
	}

	return nil
}

func NewTaskRepo(filepath, filename string) (*TaskRepo, error) {
	repo := &TaskRepo{filepath, filename}
	err := repo.init()

	return &TaskRepo{filepath, filename}, err
}

func (r *TaskRepo) Add(task *task.Task) error {
	file, err := os.OpenFile(fmt.Sprintf("%v/%v.csv", r.filepath, r.filename), os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("failed to open storage file: %v\n", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{
		strconv.Itoa(int(task.Id)),
		task.Name,
		task.Description,
		strconv.Itoa(task.Status.Int()),
		utils.HandleDateString(task.Deadline),
		utils.HandleDateString(task.DueDate),
		task.CreatedAt.Format(utils.TimeFormat),
	})
	if err != nil {
		return fmt.Errorf("failed to add item: %v\n", err)
	}

	return nil
}

func (r *TaskRepo) List() ([]*task.Task, error) {
	records, err := readCsv(r.filepath, r.filename)
	if err != nil {
		return nil, err
	}

	tasks := make([]*task.Task, 1)
	for i, record := range records {
		if i == 0 {
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("failed to convert id to int: %v\n", err)
		}

		statusInt, err := strconv.Atoi(record[3])
		if err != nil {
			return nil, fmt.Errorf("failed to convert status to int: %v\n", err)
		}

		deadline, err := utils.HandleDate(record[4])
		if err != nil {
			return nil, err
		}

		dueDate, err := utils.HandleDate(record[5])
		if err != nil {
			return nil, err
		}

		createdAt, err := utils.ParseDate(record[6])
		if err != nil {
			return nil, err
		}

		task, err := task.NewTask(
			uint16(id),
			record[1],
			record[2],
			enum.NewStatusFromInt(statusInt),
			deadline,
			dueDate,
			createdAt,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *TaskRepo) Complete(taskId uint16) error {
	records, err := readCsv(r.filepath, r.filename)
	if err != nil {
		return err
	}

	for i, record := range records {
		if i == 0 {
			continue
		}

		if record[0] == strconv.Itoa(int(taskId)) {
			records[i][3] = strconv.Itoa(enum.StatusDone)
		}
	}

	file, err := os.Create(fmt.Sprintf("%v/%v.csv", r.filepath, r.filename))
	if err != nil {
		return fmt.Errorf("failed to open file for writing: %v\n", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(records)
	if err != nil {
		return fmt.Errorf("failed to write records to file: %v\n", err)
	}

	return nil
}

func (r *TaskRepo) Delete(taskId uint16) error {
	records, err := readCsv(r.filepath, r.filename)
	if err != nil {
		return err
	}

	for i, record := range records {
		if i == 0 {
			continue
		}

		if record[0] == strconv.Itoa(int(taskId)) {
			records = append(records[:i], records[i+1:]...)
		}
	}

	file, err := os.Create(fmt.Sprintf("%v/%v.csv", r.filepath, r.filename))
	if err != nil {
		return fmt.Errorf("failed to open file for writing: %v\n", err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(records)
	if err != nil {
		return fmt.Errorf("failed to write records to file: %v\n", err)
	}

	return nil
}

func (r *TaskRepo) Exists(taskId uint16) (bool, error) {
	records, err := readCsv(r.filepath, r.filename)
	if err != nil {
		return false, err
	}

	for i, record := range records {
		if i == 0 {
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return false, fmt.Errorf("failed to convert id to int: %v\n", err)
		}

		if uint16(id) == taskId {
			return true, nil
		}
	}

	return false, nil
}

func (r *TaskRepo) GetNextId() (uint16, error) {
	records, err := readCsv(r.filepath, r.filename)
	if err != nil {
		return 0, err
	}

	lastIdx := len(records) - 1
	lastId, err := strconv.Atoi(records[lastIdx][0])
	if err != nil {
		return 0, fmt.Errorf("failed to convert id to int: %v\n", err)
	}

	return uint16(lastId + 1), nil
}

func readCsv(filepath, filename string) ([][]string, error) {
	file, err := os.Open(fmt.Sprintf("%v/%v.csv", filepath, filename))
	if err != nil {
		return nil, fmt.Errorf("failed to open storage file: %v\n", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read storage file: %v\n", err)
	}

	return records, nil
}
