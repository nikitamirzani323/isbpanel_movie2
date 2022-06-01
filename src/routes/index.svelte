<script>
    import { browser } from '$app/env';
    let tempdata = "";
    let listmovie = [
        {movie_id:1,movie_name:`Doctor Strange in the Multiverse of Madness (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/d9668308-57e5-49ee-4f2d-efa80423a500/public",movie_rate:5.6},
        {movie_id:2,movie_name:`Fantastic Beasts: The Secrets of Dumbledore (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/88ac65dc-cfc9-454e-f20e-a8ba3c09cd00/public",movie_rate:5.6},
        {movie_id:3,movie_name:`Memory (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/1ffd87fb-0ff9-4573-719f-b4c3513c2d00/public",movie_rate:5.6},
        {movie_id:4,movie_name:`Row 19 (Ryad 19) (2021)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/a29f9982-4116-4a6d-dc36-0f0f35018c00/public",movie_rate:5.6},
        {movie_id:5,movie_name:`Ambulance (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/be02eeb6-5294-403f-5744-9f697a72b900/public",movie_rate:6.3},
        {movie_id:6,movie_name:`The Girl on a Bulldozer (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/9e87b0e8-f89a-4f67-ffd2-af8e86bb9100/public",movie_rate:7.8},
        {movie_id:7,movie_name:`Top Gun: Maverick (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/74cda9ae-5675-4c60-c398-b416d7612300/public",movie_rate:8.0},
        {movie_id:8,movie_name:`Sonic the Hedgehog 2 (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/39fe7d01-c3a0-4185-8116-4f95c2810300/public",movie_rate:8.0},
        {movie_id:8,movie_name:`Gangubai Kathiawadi (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/42466eb4-7125-4a29-108b-2de1007a6900/public",movie_rate:8.0},
        {movie_id:1,movie_name:`Doctor Strange in the Multiverse of Madness (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/d9668308-57e5-49ee-4f2d-efa80423a500/public",movie_rate:5.6},
        {movie_id:2,movie_name:`Fantastic Beasts: The Secrets of Dumbledore (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/88ac65dc-cfc9-454e-f20e-a8ba3c09cd00/public",movie_rate:5.6},
        {movie_id:3,movie_name:`Memory (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/1ffd87fb-0ff9-4573-719f-b4c3513c2d00/public",movie_rate:5.6},
        {movie_id:4,movie_name:`Row 19 (Ryad 19) (2021)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/a29f9982-4116-4a6d-dc36-0f0f35018c00/public",movie_rate:5.6},
        {movie_id:5,movie_name:`Ambulance (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/be02eeb6-5294-403f-5744-9f697a72b900/public",movie_rate:6.3},
        {movie_id:6,movie_name:`The Girl on a Bulldozer (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/9e87b0e8-f89a-4f67-ffd2-af8e86bb9100/public",movie_rate:7.8},
        {movie_id:7,movie_name:`Top Gun: Maverick (2022)`,movie_img:"https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/74cda9ae-5675-4c60-c398-b416d7612300/public",movie_rate:8.0},
    ]
    async function initapp() {
        const resdata = await fetch("http://localhost:3000/api/initmovie", {
            method: "POST",
            body: JSON.stringify({
                nama: "asd",
                nik: "asd",
            }),
        });
        if (!resdata.ok) {
            const pasarMessage = `An error has occured: ${resdata.status}`;
            throw new Error(pasarMessage);
        }else{
            const jsondata = await resdata.json();
            tempdata = jsondata.data
            console.log(jsondata)
		}
	
	}
    initapp();
</script>
<div class="card w-full bg-base-300 text-neutral-content rounded-none mt-5">
    <div class="card-body p-0 mb-2">
        <h2 class="card-title border-b-2 border-pink-700 p-2">Top Rated</h2>
        <div class="grid grid-cols-2 lg:grid-cols-8 gap-2 p-2">
            {#each listmovie as rec}
                <div class="flex flex-col cursor-pointer w-full justify-items-center mb-1">
                    <div class="absolute">
                        <div class="flex mt-0 bg-black p-1 opacity-90 w-12">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-yellow-300" viewBox="0 0 20 20" fill="currentColor">
                                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                            </svg>
                            <span class=" text-sm">{rec.movie_rate}</span>
                        </div>
                    </div>
                    <img src="{rec.movie_img}" alt="{rec.movie_name}">
                    <div class="flex flex-col h-20 w-full justify-between -mt-20">
                        <span class="">&nbsp;</span>
                        <span class="w-full self-end text-center hidden lg:block lg:text-xs">{rec.movie_name}</span>
                    </div>
                </div>  
            {/each}
            
            
        </div>
    </div>
</div>
<div class="card w-full bg-base-300 text-neutral-content rounded-none mt-5">
    <div class="card-body p-0 mb-2">
        <h2 class="card-title border-b-2 border-pink-700 p-2">Popular</h2>
        <div class="grid grid-cols-2 lg:grid-cols-8 gap-2 p-2">
            {#each listmovie as rec}
                <div class="flex flex-col cursor-pointer w-full justify-items-center mb-1">
                    <div class="absolute">
                        <div class="flex mt-0 bg-black p-1 opacity-90 w-12">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-yellow-300" viewBox="0 0 20 20" fill="currentColor">
                                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                            </svg>
                            <span class=" text-sm">{rec.movie_rate}</span>
                        </div>
                    </div>
                    <img src="{rec.movie_img}" alt="{rec.movie_name}">
                    <div class="flex flex-col h-20 w-full justify-between -mt-20">
                        <span class="">&nbsp;</span>
                        <span class="w-full self-end text-center hidden lg:text-xs">{rec.movie_name}</span>
                    </div>
                </div>  
            {/each}
            
            
        </div>
    </div>
</div>
<div class="card w-full bg-base-300 text-neutral-content rounded-none mt-5">
    <div class="card-body p-0 mb-2">
        <h2 class="card-title border-b-2 border-pink-700 p-2">Terbaru</h2>
        <div class="grid grid-cols-2 lg:grid-cols-8 gap-2 p-2">
            {#each listmovie as rec}
                <div class="flex flex-col cursor-pointer w-full justify-items-center mb-1">
                    <div class="absolute">
                        <div class="flex mt-0 bg-black p-1 opacity-90 w-12">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-yellow-300" viewBox="0 0 20 20" fill="currentColor">
                                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                            </svg>
                            <span class=" text-sm">{rec.movie_rate}</span>
                        </div>
                    </div>
                    <img src="{rec.movie_img}" alt="{rec.movie_name}">
                    <div class="flex flex-col h-20 w-full justify-between -mt-20">
                        <span class="">&nbsp;</span>
                        <span class="w-full self-end text-center hidden lg:text-xs">{rec.movie_name}</span>
                    </div>
                </div>  
            {/each}
            
            
        </div>
    </div>
</div>
<div class="card w-full bg-base-300 text-neutral-content rounded-none mt-5">
    <div class="card-body p-0 mb-2">
        <h2 class="card-title border-b-2 border-pink-700 p-2">Rekomendasi</h2>
        <div class="grid grid-cols-2 lg:grid-cols-8 gap-2 p-2">
            {#each listmovie as rec}
                <div class="flex flex-col cursor-pointer w-full justify-items-center mb-1">
                    <div class="absolute">
                        <div class="flex mt-0 bg-black p-1 opacity-90 w-12">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-yellow-300" viewBox="0 0 20 20" fill="currentColor">
                                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                            </svg>
                            <span class=" text-sm">{rec.movie_rate}</span>
                        </div>
                    </div>
                    <img src="{rec.movie_img}" alt="{rec.movie_name}">
                    <div class="flex flex-col h-20 w-full justify-between -mt-20">
                        <span class="">&nbsp;</span>
                        <span class="w-full self-end text-center hidden lg:text-xs">{rec.movie_name}</span>
                    </div>
                </div>  
            {/each}
            
            
        </div>
    </div>
</div>