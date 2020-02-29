<script>
    import RouterLink from './RouterLink.svelte';
    
    const urlParams = new URLSearchParams(window.location.search);

    let relation = {datesLocations: {}}
    let artist = {}

    fetch('http://localhost:9090/api/relations?id=' + urlParams.get("id"))
        .then(res => res.json())
        .then(res => {
            console.log(res)
            relation = res
        })

    fetch('http://localhost:9090/api/artists?id=' + urlParams.get("id"))
        .then(res => res.json())
        .then(res => artist = res)
</script>

<div class="main">
    {#if artist.name}
        <table class="table">
            <tbody>
                <tr>
                    <td colspan="2">
                        <img src="{artist.image}" alt="image">
                    </td>    
                </tr>
                <tr>
                    <td>Name</td>
                    <td>{artist.name}</td>
                </tr>
                <tr>
                    <td>Members</td>
                    <td>{artist.members}</td>
                </tr>
                <tr>
                    <td>First album</td>
                    <td>{artist.firstAlbum}</td>
                </tr>
                <tr>
                    <td>Creation date</td>
                    <td>{artist.creationDate}</td>
                </tr>
                <tr>
                    <td>Creation date</td>
                    <td>{artist.creationDate}</td>
                </tr>
                {#each Object.entries(relation.datesLocations) as data}
                    <tr>
                        <td>{data[0]}</td>
                        <td>{data[1]}</td>
                    </tr>
                {/each}
            </tbody>
        </table>
    {/if}
    <RouterLink page={{path: '/', name: 'Go back'}} />

</div>

 

<style>
    .main {
        position: relative;
        min-height: 100vh;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    }
    .main-logo {
        font-size: 72px;
        font-weight: bold;
        margin-bottom: 20px;
        display: flex;
    }

    .red { color: #ea4335; }
    .blue { color: #4285f4 }
    .yellow { color: #fbbc05 }
    .green { color: #34a853 }

    .table {
        width: 600px;
        margin-bottom: 50px;
    }
    .table td:first-child {
        width: 200px;
    }
    .table td {
        padding: 10px;
    }
</style>