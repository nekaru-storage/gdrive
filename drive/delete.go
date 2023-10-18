package drive

import (
	"context"
	"fmt"
	"io"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

type DeleteArgs struct {
	Out       io.Writer
	Id        string
	Recursive bool
}

func (self *Drive) Delete(args DeleteArgs) error {
	f, err := self.service.Files.Get(args.Id).SupportsTeamDrives(true).Fields("name", "mimeType").Do()
	if err != nil {
		return fmt.Errorf("Failed to get file: %s", err)
	}

	if isDir(f) && !args.Recursive {
		return fmt.Errorf("'%s' is a directory, use the 'recursive' flag to delete directories", f.Name)
	}

	if false {
		controlledStop := fmt.Errorf("Controlled stop")
		var files []*drive.File
		pageSize := 50
		self.service.Files.List().SupportsTeamDrives(true).IncludeTeamDriveItems(true).Q(fmt.Sprintf("'%s' in parents", args.Id)).Fields([]googleapi.Field{"nextPageToken", "files(id,name)"}...).PageSize(int64(pageSize)).Pages(context.TODO(), func(fl *drive.FileList) error {
			if len(fl.Files) != 0 {
				files = append(files, fl.Files...)
			} else {
				return controlledStop
			}
			return nil
		})
		for _, f := range files {
			fmt.Printf("Deleting file: %s\n", f.Name)
			err = self.service.Files.Delete(f.Id).SupportsTeamDrives(true).Do()
			if err != nil {
				return fmt.Errorf("Failed to delete file: %s", err)
			}
		}
	}

	err = self.service.Files.Delete(args.Id).SupportsTeamDrives(true).Do()
	if err != nil {
		return fmt.Errorf("Failed to delete file: %s", err)
	}

	fmt.Fprintf(args.Out, "Deleted '%s'\n", f.Name)
	return nil
}

func (self *Drive) deleteFile(fileId string) error {
	err := self.service.Files.Delete(fileId).Do()
	if err != nil {
		return fmt.Errorf("Failed to delete file: %s", err)
	}
	return nil
}
