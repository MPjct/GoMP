package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

var Attributes_test_packets = []struct {
	packet Proto
}{
	{packet: Proto{data: StringToPacket(`
            61 03 5f 6f    73 09 64 65 62 69 61 6e    ord.a._os.debian
36 2e 30 0c 5f 63 6c 69    65 6e 74 5f 6e 61 6d 65    6.0._client_name
08 6c 69 62 6d 79 73 71    6c 04 5f 70 69 64 05 32    .libmysql._pid.2
32 33 34 34 0f 5f 63 6c    69 65 6e 74 5f 76 65 72    2344._client_ver
73 69 6f 6e 08 35 2e 36    2e 36 2d 6d 39 09 5f 70    sion.5.6.6-m9._p
6c 61 74 66 6f 72 6d 06    78 38 36 5f 36 34 03 66    latform.x86_64.f
6f 6f 03 62 61 72                                     oo.bar
`)}},
}

func Test_Attributes(t *testing.T) {
	var pkt Attributes
	var empty_bytes []byte

	// Test empty removes
	assert.Equal(t, pkt.RemoveAttribute("foobar"), false, "")
	assert.Nil(t, pkt.GetAttribute("foobar"))
	assert.Equal(t, pkt.BuildAttributes(), empty_bytes, "")

	for _, value := range Attributes_test_packets {
		pkt = value.packet.GetAttributes()
		assert.Equal(t, pkt.BuildAttributes(), value.packet.data, "")
		assert.Equal(t, pkt.RemoveAttribute("foobar"), false, "")
		assert.Equal(t, pkt.SetAttribute("foobar", "baz"), true, "")
		assert.Equal(t, pkt.SetAttribute("foobar", "baz"), true, "")
		assert.Equal(t, *pkt.GetAttribute("foobar"), "baz", "")
		assert.Equal(t, pkt.RemoveAttribute("foobar"), true, "")
	}
}
