export async function post({request,url}){
    const object =  await request.json();
    const token = object.token
    const hostname = url.host
    console.log(token)
    const resdata = await fetch("https://api.isbpanel.com/api/movie", {
        method: "POST",
        headers: {
            'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTQxNTkzOTYsIm5hbWUiOiJjVXJ7Y2wvaTpddnwxNiJ9.o4PUmAPLl3XtA291xh9S7X3c9V5abzZ-fqDogD1yCEU', 
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            "client_hostname": hostname,
        }),
    });
    const jsondata = await resdata.json();
    const status = jsondata.status;
    const message = jsondata.message;
    if(status == 200){
        return {
            status:status,
            body:{data:jsondata.record}
        }
    }else{
        return {
            status:400,
            body:{data:message}
        }
    }
    
}