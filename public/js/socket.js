function initializeSocket(socket){
    socket.onMessage(function(data) {
        if (data == "game beginning"){
            updateStatus("Game Beginning!")
        }
        else if (data == "confirmed. waiting for stranger"){
            updateStatus(data);
        }
        else{
            drawFromMsg(data);
        }

        // Echo the message back to the server.
        //socket.send("echo: " + data);
    });

    socket.on("connected", function() {
        console.log("connected");
    });

    socket.on("connecting", function() {
        console.log("connecting");
    });

    socket.on("disconnected", function() {
        console.log("disconnected");
    });

    socket.on("reconnecting", function() {
        console.log("reconnecting");
    });

    socket.on("error", function(e, msg) {
        console.log("error: " + msg);
    });

    socket.on("connect_timeout", function() {
        console.log("connect_timeout");
    });

    socket.on("timeout", function() {
        console.log("timeout");
    });

    socket.on("discard_send_buffer", function() {
        console.log("some data could not be send and was discarded.");
    });
}