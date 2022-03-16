import React, { useEffect, useState } from 'react'

export default function SSEDemo() {
    const [sse, setSSE] = useState([])
    const _sse=[]

    const append_chunk = (chunk) => {
        _sse.push(chunk)
        setSSE(_sse.slice())
    }


    // useEffect(() => {
    //     console.log("useEffect")
    //     var source =new EventSource("/api/test_sse/");

    //     source.onmessage = function (event) {
    //         console.log(event.data)
    //         append_chunk(decodeURIComponent(event.data))
    //     }
    //     // source.onerror = (e) => {
    //     //     append_chunk(`error ${JSON.stringify(e)}`)
    //     // }
    //     // source.close = (e) => {
    //     //     append_chunk(`close ${JSON.stringify(e)}`)
    //     // }
    //     // return () => {
    //     //     source.close()
    //     // }
    // }, [])

    useEffect(() => {
        // some reason sse not working properly on http-proxy-middleware

        let data=[]
        fetch('http://localhost:8000/api/test_sse/', {
            // method: 'POST',
            // headers: {
            //    'Content-Type': 'application/x-www-form-urlencoded;charset=utf-8'
            // },
            // body: querystring.stringify({ url: url })
         })      //  .then((resp) => console.log(resp))
         .catch(function (error) {
            console.log(error);
          
         })
         .then(async(response)=>{
            const reader = response.body.getReader();
            while (true) {
               const { done, value } = await reader.read();
               console.log(new TextDecoder("utf-8").decode(value))
               if (done) {
                  console.log("done",data)
                  break;
               }
               let temp_data=new TextDecoder("utf-8").decode(value).trim().split("\n\n").map(chunk=>decodeURIComponent(chunk.replace(/^data[:]/,"")))
               data=data.concat(temp_data)
               setSSE(data)               
            }               
         })
    }, [])
    


    return (
        <div>
            <h1>SSE Chunks</h1><br />
            {
                sse.map(chunk => {
                    return <p key={chunk}><b>Chunk: </b>{chunk}</p>
                })
            }
        </div>
    )
}
