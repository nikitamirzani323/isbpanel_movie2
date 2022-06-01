<script context="module">
    import { browser } from '$app/env'
    let client_token =  ""
    if(browser){
        client_token = localStorage.getItem("token");
    }
    export const load = async ({ fetch }) => {
        let listmovie = [];
        if(client_token == null){
            const res = await fetch("/api/initmovie", {
                method: "POST",
                body: JSON.stringify({}),
            })
            const record = await res.json();
            if(browser){
                localStorage.setItem("token", record.token)
            }

            
            const res2 = await fetch("/api/listmovie", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + record.token,
                },
                body: JSON.stringify({}),
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
                    token:client_token
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
</script>
<svelte:head>
    <title>ISBFILM - KOLEKSI MOVIE</title>
</svelte:head>
{#each listmovie as rec}
<div class="card w-full bg-base-300 text-neutral-content rounded-none mt-5">
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
                <img src="{rec2.movie_thumbnail}" alt="{rec2.movie_title}">
                <div class="flex flex-col h-20 w-full justify-between -mt-20">
                    <span class="">&nbsp;</span>
                    <span class="w-full self-end text-center hidden lg:block lg:text-xs">{rec2.movie_title}</span>
                </div>
            </div>  
            {/each}
        </div>
    </div>
</div>
{/each}