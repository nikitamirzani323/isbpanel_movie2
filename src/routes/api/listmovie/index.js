export async function post(event){
    const object =  await event.request.json();
    const resdata = await fetch("https://api.isbpanel.com/api/movie", {
        method: "POST",
        headers: {
            'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTQxNTkzOTYsIm5hbWUiOiJjVXJ7Y2wvaTpddnwxNiJ9.o4PUmAPLl3XtA291xh9S7X3c9V5abzZ-fqDogD1yCEU', 
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            "client_hostname": "http://localhost:7075/",
        }),
    });
    const jsondata = await resdata.json();
    console.log(jsondata)
    let listmovie = [
        {movie_id:1,movie_name:`F9 Fast & Furious 9 (2021)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/d9668308-57e5-49ee-4f2d-efa80423a500/public",movie_rate:5.6},
        {movie_id:2,movie_name:`Fantastic Beasts: The Secrets of Dumbledore (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/88ac65dc-cfc9-454e-f20e-a8ba3c09cd00/public",movie_rate:5.6},
        {movie_id:3,movie_name:`Memory (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/1ffd87fb-0ff9-4573-719f-b4c3513c2d00/public",movie_rate:5.6},
        {movie_id:4,movie_name:`Row 19 (Ryad 19) (2021)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/a29f9982-4116-4a6d-dc36-0f0f35018c00/public",movie_rate:5.6},
        {movie_id:5,movie_name:`Ambulance (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/be02eeb6-5294-403f-5744-9f697a72b900/public",movie_rate:6.3},
        {movie_id:6,movie_name:`The Girl on a Bulldozer (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/9e87b0e8-f89a-4f67-ffd2-af8e86bb9100/public",movie_rate:7.8},
        {movie_id:7,movie_name:`Top Gun: Maverick (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/74cda9ae-5675-4c60-c398-b416d7612300/public",movie_rate:8.0},
        {movie_id:8,movie_name:`Sonic the Hedgehog 2 (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/39fe7d01-c3a0-4185-8116-4f95c2810300/public",movie_rate:8.0},
        {movie_id:9,movie_name:`Gangubai Kathiawadi (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/42466eb4-7125-4a29-108b-2de1007a6900/public",movie_rate:8.0},
        {movie_id:10,movie_name:`F9 Fast & Furious 9 (2021)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/d9668308-57e5-49ee-4f2d-efa80423a500/public",movie_rate:5.6},
        {movie_id:12,movie_name:`Fantastic Beasts: The Secrets of Dumbledore (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/88ac65dc-cfc9-454e-f20e-a8ba3c09cd00/public",movie_rate:5.6},
        {movie_id:13,movie_name:`Memory (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/1ffd87fb-0ff9-4573-719f-b4c3513c2d00/public",movie_rate:5.6},
        {movie_id:14,movie_name:`Row 19 (Ryad 19) (2021)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/a29f9982-4116-4a6d-dc36-0f0f35018c00/public",movie_rate:5.6},
        {movie_id:15,movie_name:`Ambulance (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/be02eeb6-5294-403f-5744-9f697a72b900/public",movie_rate:6.3},
        {movie_id:16,movie_name:`The Girl on a Bulldozer (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/9e87b0e8-f89a-4f67-ffd2-af8e86bb9100/public",movie_rate:7.8},
        {movie_id:17,movie_name:`Top Gun: Maverick (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/74cda9ae-5675-4c60-c398-b416d7612300/public",movie_rate:8.0},
    ]
    return {
        status:200,
        body:{hello:'world',data:jsondata.record}
    }
}