import { describe, it, expect, vi } from 'vitest';
import { get } from 'svelte/store';
import { health, isEmpty, type Service } from './health';

describe('health store empty state', () => {
	it('isEmpty should be true when no services are configured', () => {
		health.set([]);
		expect(get(isEmpty)).toBe(true);
	});

	it('isEmpty should be false when services exist', () => {
		const mockService: Service = {
			name: 'Test',
			description: 'Test service',
			emoji: '🧪',
			checks: []
		};
		health.set([mockService]);
		expect(get(isEmpty)).toBe(false);
	});
});
