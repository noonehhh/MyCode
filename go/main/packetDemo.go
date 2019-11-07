package main

import (
    "fmt"
    "github.com/google/gopacket"
    "github.com/google/gopacket/layers"
    "github.com/google/gopacket/pcap"
)
/*
go语言抓包，未成功，需进一步测试优化
 */
func main(){
    fmt.Println("packet start...")
    fmt.Println("packet start...")

    deviceName := "eth0"
    snapLen := int32(65535)
    port := uint16(2121)
    filter := getFilter(port)
    fmt.Printf("device:%v, snapLen:%v, port:%v\n", deviceName, snapLen, port)
    fmt.Println("filter:", filter)

    //打开网络接口，抓取在线数据
    handle, err := pcap.OpenLive(deviceName, snapLen, true, pcap.BlockForever)
    if err != nil {
        fmt.Printf("pcap open live failed: %v", err)
        return
    }

    // 设置过滤器
    if err := handle.SetBPFFilter(filter); err != nil {
        fmt.Printf("set bpf filter failed: %v", err)
        return
    }
    defer handle.Close()

    // 抓包
    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
    packetSource.NoCopy = true
    for packet := range packetSource.Packets() {
        if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
            fmt.Println("unexpected packet")
            continue
        }

        fmt.Printf("packet:%v\n",packet)

        // tcp 层
        tcp := packet.TransportLayer().(*layers.TCP)
        fmt.Printf("tcp:%v\n", tcp)
        // tcp payload，也即是tcp传输的数据
        fmt.Printf("tcp payload:%v\n", tcp.Payload)
    }
}

func getFilter(port uint16) string {
    filter := fmt.Sprintf("tcp and ((src port %v) or (dst port %v))",  port, port)
    return filter
}