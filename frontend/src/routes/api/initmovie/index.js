export async function post({request}){
    const object =  await request.json();
    const hostname = object.hostname
    console.log(hostname)
    const resdata = await fetch("http://localhost:5000/api/initmovie", {
        method: "POST",
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            "client_hostname": hostname,
        }),
    });
    const jsondata = await resdata.json();
    const token = jsondata.token
    return {
        status:200,
        body:{token:token}
    }
}