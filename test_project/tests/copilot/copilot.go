package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func UdpNode(port int) {
	// Create a UDP socket
	conn, err := net.ListenPacket("udp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Create a buffer to read packets into
	//buf := make([]byte, 1024)

	for {
		// // Read a packet into buf
		// n, addr, err := conn.ReadFrom(buf)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// //saveMsgToDbIfNotExists("db"+port, buf[:n], buf[:n])
		// // Print the packet and sender's address
		// fmt.Printf("\nReceived on port: %d %q from %v", port, n, addr)
		ReadFromUDP(conn, port)
	}
}

func ReadFromUDP(conn net.PacketConn, port int) {
	buf := make([]byte, 1024)
	n, addr, err := conn.ReadFrom(buf)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("\nReceived %q from %v", buf[:n], addr)
	fmt.Printf("\nReceived on port: %d %q from %v\n", port, n, addr)
}

// func convertMessageToStruct(msg []byte) interface{} {
// 	var data interface{}
// 	err := json.Unmarshal(msg, &data)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return data
// }

// func convertMessageType(msg []byte) string {
// 	var data interface{}
// 	err := json.Unmarshal(msg, &data)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return data.(map[string]interface{})["type"].(string)
// }

// func setMessageType(msg []byte, msgType string) []byte {
// 	var data interface{}
// 	err := json.Unmarshal(msg, &data)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	data.(map[string]interface{})["type"] = msgType
// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return jsonData
// }

// func setMessageMap() map[string]interface{} {
// 	var data map[string]interface{}
// 	data = make(map[string]interface{})
// 	return data
// }

// func getMsgType(msg []byte) string {
// 	var data interface{}
// 	err := json.Unmarshal(msg, &data)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return data.(map[string]interface{})["type"].(string)
// }

// func setMsgType(o interface{}, t string) {
// 	o.(map[string]interface{})["type"] = t
// }

// func getYamlConfig() interface{} {
// 	yamlFile, err := ioutil.ReadFile("config.yaml")
// 	if err != nil {
// 		log.Printf("yamlFile.Get err   #%v ", err)
// 	}
// 	var data interface{}
// 	err = yaml.Unmarshal(yamlFile, &data)
// 	if err != nil {
// 		log.Fatalf("Unmarshal: %v", err)
// 	}
// 	return data //.(map[interface{}]interface{}) //(map[string]interface{})
// }

// func setYamlConfig() {
// 	config := make(map[string]interface{})
// 	config["nodes"] = []int{10000, 10001, 10002}
// 	config["node"] = 10000
// 	config["port"] = 10000
// 	config["ip"] = ""

// 	writeYamlToFileIfNotExists()
// }

// func writeYamlToFileIfNotExists() {
// 	if _, err := os.Stat("config.yaml"); os.IsNotExist(err) {
// 		yamlFile, err := os.Create("config.yaml")
// 		if err != nil {
// 			log.Printf("yamlFile.Get err   #%v ", err)
// 		}
// 		defer yamlFile.Close()
// 		yamlFile.Write([]byte("nodes:\n  - 10000\n  - 10001\n  - 10002"))
// 	}
// }

// func updateYamlConfig() {
// 	config := getYamlConfig()
// 	config["nodes"] = append(config["nodes"].([]int), 10003)
// 	config["node"] = 10003
// 	config["port"] = 10003
// 	config["ip"] = ""
// }

// func updateYamlNodes(ip string) {
// 	config := getYamlConfig()
// 	config["nodes"] = append(config["nodes"].([]int), ip)
// }

// func removeYamlNode(ip string) {
// 	config := getYamlConfig()
// 	nodes := config["nodes"].([]int)
// 	for i, node := range nodes {
// 		if node == ip {
// 			nodes = append(nodes[:i], nodes[i+1:]...)
// 		}
// 	}
// 	config["nodes"] = nodes
// }

// func getIpAddr() string {
// 	addrs, err := net.InterfaceAddrs()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, address := range addrs {
// 		// check the address type and if it is not a loopback the display it
// 		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
// 			if ipnet.IP.To4() != nil {
// 				return ipnet.IP.String()
// 			}
// 		}
// 	}
// 	return ""
// }

// func sendUdpMsgToNode(ip string, port int, msg []byte) {
// 	// Create a UDP socket
// 	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", ip, port))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer conn.Close()

// 	// Send a packet via conn
// 	_, err = conn.Write(msg)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func saveMsgToDb(db string, key []byte, value []byte) {
// 	db, err := leveldb.OpenFile(db, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	err = db.Put(key, value, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func saveMsgToDbIfNotExists(db string, key []byte, value []byte) {
// 	db, err := leveldb.OpenFile(db, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	_, err = db.Get(key, nil)
// 	if err != nil {
// 		err = db.Put(key, value, nil)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }

// func getMsgFromDb(db string, key []byte) ([]byte, error) {
// 	db, err := leveldb.OpenFile(db, nil)
// 	if err != nil {
// 		//log.Fatal(err)
// 		return nil, err
// 	}
// 	defer db.Close()

// 	data, err := db.Get(key, nil)
// 	if err != nil {
// 		//log.Fatal(err)
// 		return nil, nil
// 	}
// 	return data, nil
// }

// func testYamlConfig() {
// 	config := getYamlConfig()
// 	fmt.Println(config["nodes"])
// }

func sendUdpMsgToAllNodes(msg []byte) {
	//config := getYamlConfig()
	for _, port := range []int{10000, 10001, 10002} { //config["nodes"].([]int) {
		sendUdpMsgToNode(port, msg)
	}
}

func sendUdpMsgToNode(port int, msg []byte) {
	// Create a UDP socket
	conn, err := net.Dial("udp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Send a packet via conn
	_, err = conn.Write([]byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}

// func setMsgInCache(key string, value []byte) {
// 	cache.Set(key, value, cache.DefaultExpiration)
// }

// func getUdpMsgFromNode(port int) []byte {
// 	// Create a UDP socket
// 	conn, err := net.ListenPacket("udp", fmt.Sprintf(":%d", port))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer conn.Close()

// 	// Read a packet via conn
// 	buf := make([]byte, 1024)
// 	n, _, err := conn.ReadFrom(buf)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return buf[:n]
// }

func startSeveralNodes() {
	for i := 0; i < 3; i++ {
		go UdpNode(10000 + i)
	}
	time.Sleep(time.Millisecond * 50) //for test
}

// func startUdpNodeWithFlags() {
// 	port := flag.Int("port", 10000, "port")
// 	flag.Parse()
// 	UdpNode(*port)
// }


// func implementCacheForUdpNode() {
// 	cache := cache.New(5*time.Minute, 10*time.Minute)
// }

// func checkIfMsgInCache(key string) bool {
// 	_, found := cache.Get(key)
// 	return found
// }

// func rejectMsgIfInCache(key string) bool {
// 	if checkIfMsgInCache(key) {
// 		fmt.Println("msg in cache")
// 		return true
// 	}
// 	setMsgInCache(key, []byte{})
// 	return false
// }

func main() {
	//setYamlConfig()
	startSeveralNodes()
	//time.Sleep(time.Millisecond * 50)
	//sendUdpMsgToAllNodes([]byte("hello"))
	//sendUdpMsgToAllNodes([]byte("hello2"))
	benchmark()

}

func benchmark() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		sendUdpMsgToAllNodes([]byte("hello"))
	}
	fmt.Println(time.Since(start))
}

// func setMsgInCache(key string, value []byte) {
// 	cache.Set(key, value, cache.DefaultExpiration)
// }

// func getMsgFromCache(key string) []byte {
// 	data, found := cache.Get(key)
// 	if found {
// 		return data.([]byte)
// 	}
// 	return nil
// }

// func testExpiredMsgInCache() {
// 	setMsgInCache("key1", []byte("value1"))
// 	time.Sleep(time.Second * 6)
// 	setMsgInCache("key2", []byte("value2"))
// 	fmt.Println(getMsgFromCache("key1"))
// 	fmt.Println(getMsgFromCache("key2"))
// }
