<script>
import { curRoute } from './router.js';

export let page = {path: '/home',name: 'Home'}

function redirectTo(event){  
    curRoute.set(event.target.pathname);
    window.history.pushState(
        {path: page.path}, 
        '', 
        window.location.origin + page.path
    );
}
</script>

<a href={page.path} on:click|preventDefault={redirectTo}>{page.name}</a>

<style>
</style>