package sounds

import (
	"encoding/binary"
	"io"
	"os"

	"github.com/fatih/color"
)

func Load(buffer *[][]byte, configs map[string]string) error {
	file, err := os.Open(configs["APP_SOUND_PATH"])
	if err != nil {
		color.Red("Error opening dca file: %v", err)
		return err
	}

	var opuslen int16

	for {
		// Read opus frame length from dca file.
		err = binary.Read(file, binary.LittleEndian, &opuslen)

		// If this is the end of the file, just return.
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err := file.Close()
			if err != nil {
				return err
			}
			return nil
		}

		if err != nil {
			color.Red("Error reading from dca file :", err)
			return err
		}

		// Read encoded pcm from dca file.
		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		// Should not be any end of file errors
		if err != nil {
			color.Red("Error reading from dca file :", err)
			return err
		}

		// Append encoded pcm data to the buffer.
		*buffer = append(*buffer, InBuf)
	}
}
