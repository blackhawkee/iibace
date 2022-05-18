package main

import (
	"bufio"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/websocket-ace-inode", wsHandlerAceInode)
	http.HandleFunc("/websocket-iib-inode", wsHandlerIibInode)
	http.HandleFunc("/", indexHandler)

	fs := http.FileServer(http.Dir("static/"))

	http.HandleFunc("/static/", func(wr http.ResponseWriter, req *http.Request) {
		// Determine mime type based on the URL
		// if req.URL.RawPath.h(".css") {
		wr.Header().Set("Content-Type", "text/css")
		// }
		http.StripPrefix("/static/", fs).ServeHTTP(wr, req)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

var (
	upgrader = websocket.Upgrader{}
	tmpl     = template.Must(template.ParseFiles("templates/index-copy.htm"))
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func wsHandlerIibInode(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Continuosly read and write message
	for {
		// mt, message, err := conn.ReadMessage()
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read failed:", err)
			break
		}
		input := string(message)

		var jsonMap map[string]interface{}
		json.Unmarshal([]byte(input), &jsonMap)

		// fmt.Println(jsonMap)

		cmd := exec.Command("batch-script\\IIB_BACKUP.bat",
			jsonMap["iibNodeName"].(string),
			jsonMap["iibBackUpPath"].(string),
			jsonMap["iibInstalledPath"].(string))

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := cmd.Start(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			err := conn.WriteMessage(websocket.TextMessage, []byte(scanner.Text()))
			if err != nil {
				conn.Close()
				return
			}
		}
		if err := scanner.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := cmd.Wait(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// fmt.Println("Input message", input)

		// output := input + "\n----------------------------------------"
		// message = []byte(output)
		// err = conn.WriteMessage(mt, message)
		// if err != nil {
		// 	log.Println("write failed:", err)
		// 	break
		// }
	}

	// go func(conn *websocket.Conn) {
	// 	for {
	// 		// mType, _, err := conn.ReadMessage()
	// 		mType, msg, err := conn.ReadMessage()

	// 		fmt.Println("message", string(msg))
	// 		fmt.Println("messageType", mType)
	// 		fmt.Println("Error", strings.Split(error.Error(err), ""))

	// 		if err != nil || mType == websocket.CloseMessage {
	// 			conn.Close()
	// 			return
	// 		}
	// 	}
	// }(conn)

	// go func(conn *websocket.Conn) {

	// 	cmd := exec.Command("batch-script\\IIB_BACKUP.bat")
	// 	stdout, err := cmd.StdoutPipe()
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	if err := cmd.Start(); err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	scanner := bufio.NewScanner(stdout)
	// 	for scanner.Scan() {
	// 		err := conn.WriteMessage(websocket.TextMessage, []byte(scanner.Text()))
	// 		if err != nil {
	// 			conn.Close()
	// 			return
	// 		}
	// 	}
	// 	if err := scanner.Err(); err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	if err := cmd.Wait(); err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// }(conn)
}

func wsHandlerAceInode(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Continuosly read and write message
	for {
		// mt, message, err := conn.ReadMessage()
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read failed:", err)
			break
		}
		input := string(message)

		var jsonMap map[string]interface{}
		json.Unmarshal([]byte(input), &jsonMap)

		// fmt.Println(jsonMap)

		cmd := exec.Command("batch-script\\IIB_ACE.bat",
			jsonMap["iibBackUpNamePath"].(string),
			jsonMap["iibNodeName"].(string),
			jsonMap["aceNodeName"].(string),
			jsonMap["aceInstalledPath"].(string))

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := cmd.Start(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			err := conn.WriteMessage(websocket.TextMessage, []byte(scanner.Text()))
			if err != nil {
				conn.Close()
				return
			}
		}
		if err := scanner.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := cmd.Wait(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

	// conn, err := upgrader.Upgrade(w, r, nil)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// go func(conn *websocket.Conn) {
	// 	for {
	// 		mType, _, err := conn.ReadMessage()
	// 		if err != nil || mType == websocket.CloseMessage {
	// 			conn.Close()
	// 			return
	// 		}
	// 	}
	// }(conn)

	// go func(conn *websocket.Conn) {
	// 	cmd := exec.Command("batch-script\\IIB_ACE.bat")
	// 	stdout, err := cmd.StdoutPipe()
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	if err := cmd.Start(); err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	scanner := bufio.NewScanner(stdout)
	// 	for scanner.Scan() {
	// 		err := conn.WriteMessage(websocket.TextMessage, []byte(scanner.Text()))
	// 		if err != nil {
	// 			conn.Close()
	// 			return
	// 		}
	// 	}
	// 	if err := scanner.Err(); err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	if err := cmd.Wait(); err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// }(conn)
}
