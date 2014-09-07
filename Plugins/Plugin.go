package Plugin

import "github.com/MPjct/GoMP/MySQLProtocol"

type Plugin_interface interface {
    init(context MySQLProtocol.Context)
    read_handshake(context MySQLProtocol.Context)
    send_handshake(context MySQLProtocol.Context)
    read_auth(context MySQLProtocol.Context)
    send_auth(context MySQLProtocol.Context)
    read_auth_result(context MySQLProtocol.Context)
    send_auth_result(context MySQLProtocol.Context)
    read_query(context MySQLProtocol.Context)
    send_query(context MySQLProtocol.Context)
    read_query_result(context MySQLProtocol.Context)
    send_query_result(context MySQLProtocol.Context)
    cleanup(context MySQLProtocol.Context)
}

type Plugin struct {
}

func (plugin *Plugin) init(context MySQLProtocol.Context) {
    return
}

func (plugin *Plugin) read_handshake(context MySQLProtocol.Context) {
    return
}

func (plugin *Plugin) send_handshake(context MySQLProtocol.Context) {
    return
}

func (plugin *Plugin) read_auth(context MySQLProtocol.Context) {
    return
}

func (plugin *Plugin) send_auth(context MySQLProtocol.Context) {
    return
}

func (plugin *Plugin) read_auth_result(context MySQLProtocol.Context) {
    return
}

func (plugin *Plugin) send_auth_result(context MySQLProtocol.Context) {
    return
}

func (plugin *Plugin) read_query(context MySQLProtocol.Context) {
    return
}

func (plugin *Plugin) send_query(context MySQLProtocol.Context) {
    return
}

func (plugin *Plugin) read_query_result(context MySQLProtocol.Context) {
    return
}

func (plugin *Plugin) send_query_result(context MySQLProtocol.Context) {
    return
}

func (plugin *Plugin) cleanup(context MySQLProtocol.Context) {
    return
}
