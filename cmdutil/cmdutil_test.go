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

func TestExecCmd(t *testing.T) {
	type args struct {
		binName string
		args    []string
		workDir []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "case 1",
			args: args{
				binName: "echo",
				args:    []string{"hello"},
				workDir: nil,
			},
			want:    "hello\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExecCmd(tt.args.binName, tt.args.args, tt.args.workDir...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecCmd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExecCmd() got = %v, want %v", got, tt.want)
			}
		})
	}
}
