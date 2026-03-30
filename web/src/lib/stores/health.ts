import { writable, derived } from 'svelte/store';

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
                updateStatus: (services: Service[]) => set(services),
                fetchStatus: async () => {
                        try {
                                const response = await fetch('http://localhost:8080/api/health');
                                if (response.ok) {
                                        const data = await response.json();
                                        set(data);
                                }
                        } catch (error) {
                                console.error('Failed to fetch health status:', error);
                        }
                }
        };
}

export const health = createHealthStore();

// Derived store to check if the services list is empty
export const isEmpty = derived(health, ($health) => $health.length === 0);
