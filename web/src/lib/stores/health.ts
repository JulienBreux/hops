import { writable } from 'svelte/store';

export interface ServiceCheck {
	name: string;
	up: boolean;
	latency: number;
}

export interface Service {
	name: string;
	description: string;
	emoji: string;
	checks: ServiceCheck[];
}

function createHealthStore() {
	const { subscribe, set, update } = writable<Service[]>([]);

	return {
		subscribe,
		set,
		updateStatus: (services: Service[]) => set(services)
	};
}

export const health = createHealthStore();
