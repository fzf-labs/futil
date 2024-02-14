package cmdutil

import "testing"

func TestExecCommand(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name       string
		args       args
		wantStdout string
		wantStderr string
		wantErr    bool
	}{
		{
			name: "test1",
			args: args{
				command: "echo hello",
			},
			wantStdout: "hello",
			wantStderr: "",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStdout, gotStderr, err := ExecCommand(tt.args.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStdout != tt.wantStdout {
				t.Errorf("ExecCommand() gotStdout = %v, want %v", gotStdout, tt.wantStdout)
			}
			if gotStderr != tt.wantStderr {
				t.Errorf("ExecCommand() gotStderr = %v, want %v", gotStderr, tt.wantStderr)
			}
		})
	}
}
