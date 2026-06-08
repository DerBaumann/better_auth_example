import { env } from '$env/dynamic/private';
import { auth } from '$lib/server/auth';
import { json, type RequestHandler } from '@sveltejs/kit';
import { APIError } from 'better-auth';

const proxy: RequestHandler = async ({ request, url }) => {
	console.log('Started api request');
	let token = '';
	try {
		token = (await auth.api.getToken({ headers: request.headers })).token;
	} catch (err) {
		if (err instanceof APIError) {
			console.error('Unauthenticated');
			return json({ error: 'Unauthenticated' }, { status: 401 });
		}
	}
	console.log('token extracted');

	const targetUrl = `${env.API_URL}${url.pathname.replace('/api', '')}/${url.search}`;

	const response = await fetch(targetUrl, {
		method: request.method,
		headers: {
			...request.headers,
			Authorization: `Bearer ${token}`
		},
		body: request.method !== 'GET' ? JSON.stringify(request.body) : undefined
	});
	const data = await response.json();

	console.log('request sent and parsed');

	return json(data);
};

// handle all HTTP methods
export const GET = proxy;
export const POST = proxy;
export const PUT = proxy;
export const PATCH = proxy;
export const DELETE = proxy;
export const OPTIONS = proxy;
