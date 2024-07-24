package matrix

import (
	"bytes"
	"mime/multipart"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	EMPTY_RECORDS [][]string = make([][]string, 0)
	GOOD_RECORDS  [][]string = [][]string{{"1", "2", "3"}, {"4.0", "5.0", "6.0"}, {"7.0", "8.0", "9.0"}}
	BAD_RECORDS   [][]string = [][]string{{"A", "B", "C"}, {"4.0", "5.0", "6.0"}, {"7.0", "8.0", "9.0"}}
)

func TestMatrixProduct(t *testing.T) {
	t.Run("matrix product: positive test case 1", func(t *testing.T) {
		want := 362880.0
		matrixService := matrixServiceImpl{}
		got, err := matrixService.matrixProduct(GOOD_RECORDS)
		require.Nil(t, err)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("matrix product: positive test case 2", func(t *testing.T) {
		want := 0.0
		matrixService := matrixServiceImpl{}
		got, err := matrixService.matrixProduct(EMPTY_RECORDS)
		require.Nil(t, err, "error should be nil")
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("matrix product: negative test case 1", func(t *testing.T) {
		want := 0.0
		matrixService := matrixServiceImpl{}
		got, err := matrixService.matrixProduct(BAD_RECORDS)
		require.NotNil(t, err, "based on records args passed err should not be nil")
		require.Equal(t, got, want, "based on records args passed result should be nil")
	})
}

func TestMatrixSum(t *testing.T) {
	t.Run("matrix sum: positive test case 1", func(t *testing.T) {
		want := 45
		matrixService := matrixServiceImpl{}
		got, err := matrixService.matrixSum(GOOD_RECORDS)
		require.Nil(t, err, "error should be nil")
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("matrix sum: positive test case 2", func(t *testing.T) {
		want := 0
		matrixService := matrixServiceImpl{}
		got, err := matrixService.matrixSum(EMPTY_RECORDS)
		require.Nil(t, err, "error should be nil")
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("matrix sum: negative test case 1", func(t *testing.T) {
		want := 0
		matrixService := matrixServiceImpl{}
		got, err := matrixService.matrixSum(BAD_RECORDS)
		require.NotNil(t, err, "based on records args passed err should not be nil")
		require.Equal(t, got, want, "based on records args passed result should be 0")
	})
}

func TestFlattenMatrix(t *testing.T) {
	t.Run("flatten matrix: positive test case 1", func(t *testing.T) {
		want := "1,2,3,4.0,5.0,6.0,7.0,8.0,9.0"
		matrixService := matrixServiceImpl{}
		got := matrixService.flattenMatrix(GOOD_RECORDS)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("flatten matrix: positive test case 2", func(t *testing.T) {
		want := ""
		matrixService := matrixServiceImpl{}
		got := matrixService.flattenMatrix(EMPTY_RECORDS)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestInvertMatrix(t *testing.T) {
	t.Run("invert matrix: positive test case 1", func(t *testing.T) {
		want := make([][]string, 0)
		matrixService := matrixServiceImpl{}
		got := matrixService.invertMatrix(EMPTY_RECORDS)
		if len(got) != len(want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("invert matrix: positive test case 2", func(t *testing.T) {
		want := [][]string{{"1", "4.0", "7.0"}, {"2", "5.0", "8.0"}, {"3", "6.0", "9.0"}}
		matrixService := matrixServiceImpl{}
		got := matrixService.invertMatrix(GOOD_RECORDS)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestFormatMatrix(t *testing.T) {
	t.Run("invert matrix: positive test case 1", func(t *testing.T) {
		want := ""
		matrixService := matrixServiceImpl{}
		got := matrixService.formatMatrix(EMPTY_RECORDS)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("invert matrix: positive test case 2", func(t *testing.T) {
		want := "1,2,3\n4.0,5.0,6.0\n7.0,8.0,9.0\n"
		matrixService := matrixServiceImpl{}
		got := matrixService.formatMatrix(GOOD_RECORDS)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}



// MockFileHeader mocks multipart.FileHeader

type MockFile struct {
	*bytes.Reader
}

func (m MockFile) Close() error {
	return nil
}


type MockFileHeader struct {
	mockOpen func() (multipart.File, error)
}

func (m *MockFileHeader) Open() (multipart.File, error) {
	return m.mockOpen()
}

func mockFileHeader(filename string, content string) *multipart.FileHeader {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filename)
	part.Write([]byte(content))
	writer.Close()

	reader := multipart.NewReader(body, writer.Boundary())
	form, _ := reader.ReadForm(int64(body.Len()))
	return form.File["file"][0]
}


func TestParseCsvFile(t *testing.T) {
	ms := &matrixServiceImpl{}

	tests := []struct {
		name        string
		fileContent string
		expected    [][]string
		expectError bool
	}{
		{
			name:        "Valid CSV",
			fileContent: "a,b,c\n1,2,3\n4,5,6",
			expected:    [][]string{{"a", "b", "c"}, {"1", "2", "3"}, {"4", "5", "6"}},
			expectError: false,
		},
		{
			name:        "Empty CSV",
			fileContent: "",
			expected:    make([][]string, 0),
			expectError: false,
		},
		{
			name:        "Invalid CSV (Unbalanced Quotes)",
			fileContent: "a,\"b,c\n1,2,3",
			expected:    make([][]string, 0),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := mockFileHeader("test.csv", tt.fileContent)
			result, err := ms.parseCsvFile(file)

			if tt.expectError && err == nil {
				t.Errorf("Expected an error, but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if len(result) != len(tt.expected) && len(result) == 0 {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}
