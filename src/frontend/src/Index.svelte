<script>
    import Nested from './nested.svelte';
    let name = 'world';
    let htmlString = `this string contains some <strong>HTML!!!</strong>`;
    let src = ''
</script>

<style>
	p {
		color: purple;
		font-family: 'Comic Sans MS', cursive;
		font-size: 2em;
	}

  h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}
</style>

<h1>Hello {name.toUpperCase()}!</h1>
<p>{@html htmlString}</p>
<img {src} alt="A man dances.">
<Nested></Nested>