function parseProfile(body) {
  const value = JSON.parse(body);
  if (typeof value.name !== 'string' || value.name.trim() === '') {
    throw new TypeError('name is required');
  }
  return { name: value.name.trim() };
}

export function profileEndpoint(requestBody) {
  try {
    const profile = parseProfile(requestBody);
    return {
      status: 201,
      headers: { 'content-type': 'application/json' },
      body: JSON.stringify(profile),
    };
  } catch {
    return {
      status: 400,
      headers: { 'content-type': 'application/json' },
      body: JSON.stringify({ error: 'invalid profile' }),
    };
  }
}
