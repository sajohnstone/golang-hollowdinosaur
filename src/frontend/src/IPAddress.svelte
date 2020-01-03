<script>
    let query = "";
    let result;

    async function getResult() {
      console.log("Fetch results..")
      let response = await fetch(`http://localhost:3000/api/whoami`);
      let text = await response.json();
      console.log(text);
      let body = text.message;
      return body;
    }

    function handleClick(e) {
      result = getResult();
    }
</script>

<style>
  .sitetitle {
    color: #f35626;
    background-image: linear-gradient(92deg, #f35626 0%,#feab3a 100%);
    -webkit-background-clip: text;
            background-clip: text;
    -webkit-text-fill-color: transparent;
    font-size: 3rem;
    line-height: 1;
  }
</style>

{#if result===undefined}
  <h1 class="sitetitle"> </h1>
  {:else}
  {#await result}
  <h1 class="sitetitle">Loading...</h1>
  {:then value}
  <h1 class="animated infinite pulse sitetitle">{value}</h1>
  {:catch error}
  <p>{error.message}</p>
  {/await}
{/if}

<button on:click={handleClick} class="pure-button pure-button-primary">Show</button>