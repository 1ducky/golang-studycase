package csv

import (
	"context"
	"encoding/csv"
	"io"
	"log"
)

func ReadCsvFile(ctx context.Context, reader io.Reader) <-chan []string {
	RowStream := make(chan []string)
	go func() {
		defer close(RowStream)
		reader := csv.NewReader(reader)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				for {
					record, err := reader.Read()
					if err == io.EOF {
						return
					}
					if err != nil {
						log.Printf("error when reading csv : %v\n", err)
						continue
					}
					RowStream <- record
				}
			}
		}
	}()
	return RowStream
}
