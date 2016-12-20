function initClient(socket){
    socket.onmessage = onmessageHandler;
    socket.onopen = onopenHandler;
    socket.onclose = oncloseHandler;
    socket.onerror = onerrorHandler;
    return {
        socket: socket
    }

}

function onmessageHandler(event){
    console.log("message received: ");
    console.log(event.data);
}

function onopenHandler(){
    console.log("connection open");
}

function oncloseHandler(){
    console.log("connection closed");
}

function onerrorHandler(event){
    console.log("error");
}