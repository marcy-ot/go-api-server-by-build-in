package repository

import (
	"log"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/marcy-ot/go-api-server-by-build-in/domain"
)

var todoBackupFilePath = "../data_store/todo_bk.json"

func TestMain(m *testing.M) {
	// Resolve the file path
	todoFilePath = "../data_store/todo.json"

	// テストを実行
	code := m.Run()

	// テスト終了後、todo_bk.jsonの内容をtodo.jsonにコピー
	backupData, err := os.ReadFile(todoBackupFilePath)
	if err != nil {
		log.Fatalf("バックアップファイルの読み込みに失敗しました: %v", err)
	}

	err = os.WriteFile(todoFilePath, backupData, 0644)
	if err != nil {
		log.Fatalf("todo.jsonの復元に失敗しました: %v", err)
	}

	os.Exit(code)

}

func Test_ReadTodo(t *testing.T) {
	expected := []domain.Todo{
		{
			Id:        1,
			Title:     "サンプルタスク",
			Content:   "これはサンプルのタスク内容です",
			CreatedAt: time.Date(2024, 3, 20, 10, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2024, 3, 20, 10, 0, 0, 0, time.UTC),
		},
	}

	todo := ReadTodo()
	if !reflect.DeepEqual(todo, expected) {
		t.Errorf("FetchTodo() = %v, want %v", todo, expected)
	}
}

func Test_SearchTodo(t *testing.T) {
	tests := []struct {
		name     string
		want     []domain.Todo
		arg      domain.TodoSearchCondition
		hasError bool
	}{
		{
			name: "search todo",
			want: []domain.Todo{{Id: 1,
				Title:     "サンプルタスク",
				Content:   "これはサンプルのタスク内容です",
				CreatedAt: time.Date(2024, 3, 20, 10, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2024, 3, 20, 10, 0, 0, 0, time.UTC)},
			},
			arg:      domain.TodoSearchCondition{Title: "サンプルタスク"},
			hasError: false,
		},
		{
			name:     "todo not found",
			want:     nil,
			arg:      domain.TodoSearchCondition{Title: "unknown"},
			hasError: false,
		},
	}

	for _, tt := range tests {
		actual := SearchTodo(tt.arg)

		if !reflect.DeepEqual(actual, tt.want) {
			t.Errorf("FindTodo() = %v, want %v", actual, tt.want)
		}
	}
}

func Test_FindTodo(t *testing.T) {
	tests := []struct {
		name     string
		want     domain.Todo
		arg      int
		hasError bool
	}{
		{
			name: "find todo",
			want: domain.Todo{
				Id:        1,
				Title:     "サンプルタスク",
				Content:   "これはサンプルのタスク内容です",
				CreatedAt: time.Date(2024, 3, 20, 10, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2024, 3, 20, 10, 0, 0, 0, time.UTC),
			},
			arg:      1,
			hasError: false,
		},
		{
			name:     "todo not found",
			want:     domain.Todo{},
			arg:      999,
			hasError: true,
		},
	}

	for _, tt := range tests {
		actual, err := FindTodo(tt.arg)
		if err != nil {
			if !tt.hasError {
				t.Errorf("FindTodo() = %v, want %v", actual, tt.want)
			}
			return
		}

		if !reflect.DeepEqual(actual, tt.want) {
			t.Errorf("FindTodo() = %v, want %v", actual, tt.want)
		}
	}
}

func Test_ResisterTodo(t *testing.T) {
	want := []domain.Todo{
		{
			Id:        1,
			Title:     "サンプルタスク",
			Content:   "これはサンプルのタスク内容です",
			CreatedAt: time.Date(2024, 3, 20, 10, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2024, 3, 20, 10, 0, 0, 0, time.UTC),
		},
		{
			Id:        2,
			Title:     "サンプルタスク2",
			Content:   "これはサンプルのタスク内容です2",
			CreatedAt: time.Date(2024, 3, 20, 10, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2024, 3, 20, 10, 0, 0, 0, time.UTC),
		},
	}

	arg := domain.Todo{
		Title:     "サンプルタスク2",
		Content:   "これはサンプルのタスク内容です2",
		CreatedAt: time.Date(2024, 3, 20, 10, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2024, 3, 20, 10, 0, 0, 0, time.UTC),
	}

	err := ResisterTodo(arg)
	if err != nil {
		t.Errorf("ResisterTodo() = %v, want %v", err, nil)
	}

	actual := ReadTodo()
	if !reflect.DeepEqual(actual, want) {
		t.Errorf("ResisterTodo() = %v, want %v", actual, want)
	}
}
