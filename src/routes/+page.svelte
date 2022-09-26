<script lang="ts" >
import axios from "axios"

const shortenedUrl = "https:fdgdfdgf"
const apiUrl = "http://localhost:3000/"
let copied = false
let loading = false
$: isLoading = loading
$: isCopied = copied
let longUrl = ""

const shortenUrl = async (url:string) => {
    const data = {
        "url": longUrl
    }
    loading = true
    const result = await axios.post(
        `${apiUrl}shorten`,
        data
    ).then(
        res => res.data
    )
    alert(result)
}

function copyToClipboard(value: string){
    navigator.clipboard.writeText(value)
    copied = true
}
</script>

<h1 class=" bg-green-500 flex justify-around  font-mono text-lg md:text-xl lg:text-2xl text-white">
    <div>Go Go Short</div>
    <button>
        Login/logout/Account
    </button>
</h1>
<div class=" flex justify-center gap-4 my-4">
    <input bind:value="{longUrl}" class=" p-3 border-green-500 border-2 rounded" type="text" placeholder="Enter your link here">
    <button on:click="{()=> shortenUrl(longUrl)}"  class=" bg-green-500 text-white px-2 rounded">
        {#if isLoading}
            <p>Shorten</p>
        {:else}
            <p class=" animate-bounce text-3xl font-extrabold">.</p>
        {/if}
    </button>
</div>
<div class=" bg-slate-600 flex justify-center gap-5">
<div class="  text-white px-3" >{shortenedUrl}</div>
<button  on:click="{()=> copyToClipboard(shortenedUrl)}">
    {#if isCopied}
        <p class=" text-green-400">Copied</p>
    {:else}
        <p class=" text-white">Copy</p>
        
    {/if}
</button>
</div>