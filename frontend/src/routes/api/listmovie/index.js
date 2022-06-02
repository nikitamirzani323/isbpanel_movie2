export async function post({request,url}){
    const object =  await request.json();
    const token = object.token
    const hostname = object.hostname
    const resdata = await fetch("http://localhost:5000/api/listmovie", {
        method: "POST",
        headers: {
            Authorization: "Bearer " + token,
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