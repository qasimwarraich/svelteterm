    export async function SpawnInstance() {
        const response = await fetch('https://0.0.0.0:8080/api/createcontainer');
        const container = await response.json();
        await StartContainer(container);
        return container
    }

    export async function StartContainer(container) {
        let options = {
            method: 'Post',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(container),
        };
        fetch('https://0.0.0.0:8080/api/startcontainer', options);
    }

    export async function StopContainer(container) {
        let options = {
            method: 'Post',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(container),
        };
        fetch('https://0.0.0.0:8080/api/stopcontainer', options);
    }

    export async function ResizeContainerPTY(resize) {
        let options = {
            method: 'Post',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(resize),
        };
        fetch('https://0.0.0.0:8080/api/resizecontainer', options);
    }
