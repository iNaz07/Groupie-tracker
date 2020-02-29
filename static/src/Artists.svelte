<script>
    import RouterLink from './RouterLink.svelte';

    let store = []
    let artists = []

    function fetchLocaations() {
        fetch('http://localhost:9090/api/locations')
            .then(res => res.json())
            .then(res => {
                if (res && res.index) {
                    res.index.forEach(item => {
                        for (let i = 0; i < store.length; i++) {
                            if (store[i].id === item.id) {
                                store[i].locations = item.locations
                                break;
                            }
                        }
                    });
                }
            })
    }

    fetch('http://localhost:9090/api/artists')
        .then(res => res.json())
        .then(res => {
            store = res
            artists = res
            fetchLocaations()
        })

    

    function filter(value) {
        artists = store
        artists = artists.map(artist => {
            if (value === "") {
                artist.founded = []
                return artist
            }

            let founded = false
            artist.founded = []
            if (artist.name.toLowerCase().includes(value.toLowerCase())) {
                founded = true
                artist.founded.push('name')
            } 
            if (artist.members.toString().toLowerCase().includes(value.toLowerCase())) {
                founded = true
                artist.founded.push('member')
            } 
            if (artist.firstAlbum.includes(value)) {
                founded = true
                artist.founded.push('first album')
            }
            if (String(artist.creationDate).includes(value)) {
                founded = true
                artist.founded.push('creation date')
            }
            if (artist.locations.toString().toLowerCase().includes(value.toLowerCase())) {
                founded = true
                artist.founded.push('locations')
            }

            if (founded) {
                return artist
            }
        }).filter(item => item != undefined)

    }
</script>

<div>
    <header>
        <div class="header-logo">
            <span class="blue">G</span>
            <span class="red">r</span>
            <span class="yellow">o</span>
            <span class="green">u</span>
            <span class="blue">p</span>
            <span class="red">i</span>
            <span class="yellow">e &nbsp;</span>
            <span class="green">T</span>
            <span class="red">r</span>
            <span class="yellow">a</span>
            <span class="green">c</span>
            <span class="blue">k</span>
            <span class="yellow">e</span>
            <span class="red">r</span>
        </div>

        <div class="input-wrapper">
            <svg focusable="false" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                <path d="M15.5 14h-.79l-.28-.27A6.471 6.471 0 0 0 16 9.5 6.5 6.5 0 1 0 9.5 16c1.61 0 3.09-.59 4.23-1.57l.27.28v.79l5 4.99L20.49 19l-4.99-5zm-6 0C7.01 14 5 11.99 5 9.5S7.01 5 9.5 5 14 7.01 14 9.5 11.99 14 9.5 14z"/>
            </svg>
            <input on:input="{e => filter(e.target.value)}">
        </div>
    </header>

    <div class="cards">
        {#each artists as { id, name, image, founded }}
            <div class="card">
                <img src="{image}" alt="Image">
                <div class="flex">
                    <h4>{name}</h4>
                    <RouterLink page={{path: '/info?id=' + id, name: 'more'}} />
                </div>
                {#if founded}
                    <div class="card-founded">
                        {#each founded as foundCategory}
                            <span>{foundCategory} &nbsp;</span>
                        {/each}
                    </div>
                {/if}
            </div>
        {/each}
    </div>
</div>

<style>
    * {
        box-sizing: border-box
    }
    header {
        margin-bottom: 50px;
        display: flex;
        align-items: center;
    }
    .header-logo {
        padding: 10px;
        font-size: 18px;
        display: flex;
        font-weight: bold;
    }
    .red { color: #ea4335; }
    .blue { color: #4285f4 }
    .yellow { color: #fbbc05 }
    .green { color: #34a853 }
    .flex { display: flex; align-items: center; justify-content: space-between; }

    .cards {
        display: flex;
        flex-wrap: wrap;
        padding: 0 50px;
    }
    .card {
        border: 1px solid #dfe1e5;
        width: calc(15% - 20px);
        min-width: 200px;
        border-radius: 8px;
        padding: 20px;
        margin-bottom: 20px;
        margin-right: 5px;
    }
    .card:hover {
        box-shadow: 0 1px 6px 0 rgba(32,33,36,0.28);
        border-color: rgba(223,225,229,0);
    }
    .card img {
        width: 100%;
    }
    .card-founded {
        font-size: 10px;
        color: #777;
    }

    .input-wrapper {
        display: flex;
        align-items: center;
        height: 38px;
        border: 1px solid #dfe1e5;
        width: 580px;
        padding: 0 15px;
        border-radius: 23px;
    }
    .input-wrapper:hover {
        box-shadow: 0 1px 6px 0 rgba(32,33,36,0.28);
        border-color: rgba(223,225,229,0);
    }
    .input-wrapper svg {
        width: 24px;
        height: 24px;
        fill: #9AA0A6;
        margin-right: 5px;
    }
    .input-wrapper input {
        width: 100%;
        height: 36px;
        border: none;
        font-size: 14px;
        margin: 0;
        padding: 0;
    }

</style>