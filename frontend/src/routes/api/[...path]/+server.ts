import { env } from '$env/dynamic/private';
import { auth } from '$lib/server/auth';
import { json, type RequestHandler } from '@sveltejs/kit';
import { APIError } from 'better-auth';

const proxy: RequestHandler = async ({ request, url }) => {
	let token = '';
	try {
		token = (await auth.api.getToken({ headers: request.headers })).token;
	} catch (err) {
		if (err instanceof APIError) {
			return json({ error: 'Unauthenticated' }, { status: 401 });
		}
	}

	const targetUrl = `${env.API_URL}${url.pathname.replace('/api', '')}/${url.search}`;
	const body = await request.json();

	const response = await fetch(targetUrl, {
		method: request.method,
		headers: {
			...request.headers,
			Authorization: `Bearer ${token}`
		},
		body: request.method !== 'GET' ? JSON.stringify(body) : undefined
	});
	const data = await response.json();

	return json(data);
};

// handle all HTTP methods
export const GET = proxy;
export const POST = proxy;
export const PUT = proxy;
export const PATCH = proxy;
export const DELETE = proxy;
export const OPTIONS = proxy;
