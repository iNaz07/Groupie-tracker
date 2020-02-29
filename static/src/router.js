import Home from './Home.svelte';
import Artists from './Artists.svelte';

import { writable } from 'svelte/store';

const router = {
    '/': Artists,
    '/info': Home
}

export default router;
export const curRoute = writable('/home');