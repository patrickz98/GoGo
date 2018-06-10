window.addEventListener("load", function(evt)
{
    let output = document.getElementById("output");
    let input = document.getElementById("input");
    let ws;

    let print = function(message)
    {
        let d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };

    document.getElementById("open").onclick = function(evt)
    {
        if (ws)
        {
            return false;
        }

        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt)
        {
            print("OPEN");
        };

        ws.onclose = function(evt)
        {
            print("CLOSE");
            ws = null;
        };

        ws.onmessage = function(evt)
        {
            print("RESPONSE: " + evt.data);
        };

        ws.onerror = function(evt)
        {
            print("ERROR: " + evt.data);
        };

        return false;
    };

    document.getElementById("send").onclick = function(evt)
    {
        if (! ws)
        {
            return false;
        }

        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };

    document.getElementById("close").onclick = function(evt)
    {
        if (! ws)
        {
            return false;
        }

        ws.close();
        return false;
    };
});
