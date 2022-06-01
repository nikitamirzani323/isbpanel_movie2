export async function post(event){
    const object =  await event.request.json();
    // const nama = object.nama;
    // const nik = object.nik;
    // const header = event.request;
    // console.log(header)
    const resdata = await fetch("https://api.isbpanel.com/api/init", {
        method: "POST",
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            "client_hostname": "http://localhost:3000/",
        }),
    });
    const jsondata = await resdata.json();
    const token = jsondata.token
    return {
        status:200,
        body:{hello:'world',token:token}
    }
}