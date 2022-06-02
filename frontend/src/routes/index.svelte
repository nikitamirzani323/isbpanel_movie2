<script context="module">
    import { browser } from '$app/env'
    let client_token =  ""
    if(browser){
        client_token = localStorage.getItem("token");
    }
    export const load = async ({ fetch,url }) => {
        let listmovie = [];
        if(client_token == null){
            const res = await fetch("/api/initmovie", {
                method: "POST",
                body: JSON.stringify({
                    hostname:url.host
                }),
            })
            const record = await res.json();
            if(browser){
                localStorage.setItem("token", record.token)
            }

            
            const res2 = await fetch("/api/listmovie", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    token:record.token
                }),
            })
            const record2 = await res2.json();
            listmovie = record2.data
            return {
                props: {
                    listmovie,
                }
            }
        }else{
            console.log(client_token)
            const res = await fetch("/api/listmovie", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    token:client_token,
                    hostname:url.host
                }),
            })
            const record = await res.json();
            listmovie = record.data
            return {
                props: {
                    listmovie,
                }
            } 
        }  
        
    };
</script>
<script>
    export let listmovie;
    const loaded = new Map();
    function lazy(node, data) {
		if (loaded.has(data.src)) {
			node.setAttribute('src', data.src);
		} else {
			// simulate slow loading network
			setTimeout(() => {
				const img = new Image();
				img.src = data.src;
				img.onload = () => {
					loaded.set(data.src, img);
					node.setAttribute('src', data.src); 
				};
			}, 100);
		}
		return {
			destroy(){} // noop
		};
	}
</script>
<svelte:head>
    <title>ISBFILM - KOLEKSI MOVIE</title>
</svelte:head>
{#each listmovie as rec}
<section class="card w-full bg-base-300 text-neutral-content rounded-none mt-5">
    <div class="card-body p-0 mb-2">
        <h2 class="card-title border-b-2 border-pink-700 p-2">{rec.movie_category} </h2>
        <div class="grid grid-cols-2 lg:grid-cols-8 gap-2 p-2">
            {#each rec.movie_list as rec2}
            <div class="flex flex-col cursor-pointer w-full justify-items-center mb-1">
                <div class="absolute">
                    <div class="flex mt-0 bg-black p-1 opacity-90 w-12">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-yellow-300" viewBox="0 0 20 20" fill="currentColor">
                            <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                        </svg>
                        <span class=" text-sm">5.2</span>
                    </div>
                </div>
                <img class="w-full"
                    use:lazy="{{src: rec2.movie_thumbnail}}" 
                    src="https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/fd0287a2-353d-4b47-9a6c-9c8df2ab3f00/public" alt="{rec2.movie_title}" >
                <div class="flex flex-col h-20 w-full justify-between -mt-20">
                    <span class="">&nbsp;</span>
                    <span class="w-full self-end text-center hidden lg:block lg:text-xs">{rec2.movie_title}</span>
                </div>
            </div>  
            {/each}
        </div>
    </div>
</section>
{/each}