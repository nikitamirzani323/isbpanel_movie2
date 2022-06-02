import { MY_GO_BACKEND_URL } from '$lib/Env';
export async function post({request}){
    const path_api = MY_GO_BACKEND_URL
    const object =  await request.json();
    const hostname = object.hostname
    console.log(hostname)
    const resdata = await fetch(path_api+"/api/initmovie", {
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