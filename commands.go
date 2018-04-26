package flexvolume

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli"
)

func Commands(fv FlexVolume) []cli.Command {
	return []cli.Command{
		{
			Name:  "init",
			Usage: "Initialize the driver",
			Action: func(c *cli.Context) error {
				return handle(fv.Init())
			},
		},
		{
			Name:  "getvolumename",
			Usage: "Returns the unique name of the volume",
			Action: func(c *cli.Context) error {
				var opts map[string]string

				if err := json.Unmarshal([]byte(c.Args().Get(1)), &opts); err != nil {
					return err
				}

				return handle(fv.GetVolumeName(opts))
			},
		},
		{
			Name:  "attach",
			Usage: "Attach the volume",
			Action: func(c *cli.Context) error {
				var opts map[string]string

				if err := json.Unmarshal([]byte(c.Args().Get(1)), &opts); err != nil {
					return err
				}

				return handle(fv.Attach(opts))
			},
		},
		{
			Name:  "waitforattach",
			Usage: "Waits until a volume is fully attached to a node and its device emerges",
			Action: func(c *cli.Context) error {
				var opts map[string]string

				if err := json.Unmarshal([]byte(c.Args().Get(2)), &opts); err != nil {
					return err
				}

				return handle(fv.WaitForAttach(c.Args().Get(1), opts))
			},
		},
		{
			Name:  "detach",
			Usage: "Detach the volume",
			Action: func(c *cli.Context) error {
				return handle(fv.Detach(c.Args().Get(1), c.Args().Get(2)))
			},
		},
		{
			Name:  "isattached",
			Usage: "Checks that a volume is attached to a node",
			Action: func(c *cli.Context) error {
				var opts map[string]string

				if err := json.Unmarshal([]byte(c.Args().Get(1)), &opts); err != nil {
					return err
				}

				return handle(fv.Detach(opts, c.Args().Get(2)))
			},
		},
		{
			Name:  "mountdevice",
			Usage: "Mounts a volume’s device to a directory",
			Action: func(c *cli.Context) error {
				var opts map[string]string

				if err := json.Unmarshal([]byte(c.Args().Get(3)), &opts); err != nil {
					return err
				}

				return handle(fv.MountDevice(c.Args().Get(1), c.Args().Get(2), opts))
			},
		},
		{
			Name:  "unmountdevice",
			Usage: "Unmounts a volume’s device from a directory",
			Action: func(c *cli.Context) error {
				return handle(fv.UnmountDevice(c.Args().Get(1)))
			},
		},
		{
			Name:  "mount",
			Usage: "Mount the volume",
			Action: func(c *cli.Context) error {
				var opts map[string]string

				if err := json.Unmarshal([]byte(c.Args().Get(1)), &opts); err != nil {
					return err
				}

				return handle(fv.Mount(c.Args().Get(1), c.Args().Get(2), opts))
			},
		},
		{
			Name:  "umount",
			Usage: "Mount the volume",
			Action: func(c *cli.Context) error {
				return handle(fv.Unmount(c.Args().Get(1)))
			},
		},
	}
}

// The following handles:
//   * Output of the Response object.
//   * Sets an error so we can bubble up an error code.
func handle(resp Response) error {
	// Format the output as JSON.
	output, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}
