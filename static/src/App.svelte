<script>
	import router, { curRoute } from './router.js';
	import RouterLink from './RouterLink.svelte';
	import { onMount } from 'svelte';
	
	onMount(() => {  
		curRoute.set(window.location.pathname);  

		if (!history.state) {    
			window.history.replaceState({path: window.location.pathname}, '',   window.location.href)  
		}
	})
		
	function handlerBackNavigation(event){  
		curRoute.set(event.state.path)
	}


</script>

<svelte:window on:popstate={handlerBackNavigation} />

<div id="pageContent">
<svelte:component this={router[$curRoute]} />
</div> 
