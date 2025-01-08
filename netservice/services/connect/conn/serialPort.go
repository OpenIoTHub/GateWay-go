package conn

import (
	"github.com/OpenIoTHub/utils/io"
	"github.com/OpenIoTHub/utils/models"
	"github.com/jacobsa/go-serial/serial"
	"log"
	"net"
)

func JoinSerialPort(stream net.Conn, m *models.ConnectSerialPort) error {
	options := serial.OpenOptions(*m)
	conn, err := serial.Open(options)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	go io.Join(stream, conn)
	return nil
}
