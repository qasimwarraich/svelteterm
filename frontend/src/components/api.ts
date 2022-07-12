import type { Container, ResizeRequest } from '../types/types';

export async function SpawnInstance(): Promise<Container> {
    const response = await fetch(import.meta.env.VITE_APIENDPOINT + '/api/createcontainer');
    const container = await response.json();
    await StartContainer(container);
    return container;
}

export async function StartContainer(container: Container) {
    let options = {
        method: 'Post',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(container),
    };
    fetch(import.meta.env.VITE_APIENDPOINT + '/api/startcontainer', options);
}

export async function StopContainer(container: Container) {
    let options = {
        method: 'Post',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(container),
    };
    fetch(import.meta.env.VITE_APIENDPOINT + '/api/stopcontainer', options);
}

export async function ResizeContainerPTY(resize: ResizeRequest) {
    let options = {
        method: 'Post',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(resize),
    };
    fetch(import.meta.env.VITE_APIENDPOINT + '/api/resizecontainer', options);
}
