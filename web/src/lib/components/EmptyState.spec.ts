/// <reference types="@testing-library/jest-dom" />
import { describe, it, expect, afterEach } from 'vitest';
import { render, cleanup } from '@testing-library/svelte';
import EmptyState from './EmptyState.svelte';

describe('EmptyState component', () => {
	afterEach(() => {
		cleanup();
	});

	it('should render the illustration and quick start guide', () => {
		const { getByText, container } = render(EmptyState);

		expect(getByText('No configured service.')).toBeInTheDocument();
		expect(getByText('Quick Start Guide')).toBeInTheDocument();
		expect(container.querySelector('svg')).toBeInTheDocument(); // Illustration check
	});

	it('should display a YAML configuration example', () => {
		const { getByText } = render(EmptyState);

		// A snippet of the expected code block should be present
		expect(getByText(/services:/)).toBeInTheDocument();
		expect(getByText(/CONFIG_PATH/)).toBeInTheDocument();
	});
});
