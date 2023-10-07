import {
  createCookieSessionStorage,
  createFileSessionStorage,
  createMemorySessionStorage,
  redirect,
} from '@remix-run/node';
import type { SessionIdStorageStrategy } from '@remix-run/node';
import invariant from 'tiny-invariant';
import type { NhostSession } from '@nhost/nhost-js';
// import type { User } from '~/models/user.server';
// import { getUserById } from '~/models/user.server';

invariant(process.env.SESSION_SECRET, 'SESSION_SECRET must be set');

export type ZberatelSession = {
  auth?: NhostSession;
};

const secrets = [process.env.SESSION_SECRET];

const getCreateSessiontStorage = () => {
  const cookie: SessionIdStorageStrategy['cookie'] = {
    name: '__session' as const,
    httpOnly: true as const,
    path: '/',
    sameSite: 'lax',
    secrets,
    secure: process.env.NODE_ENV === 'production',
  };
  if (process.env.SESSION_STORAGE === 'cookie') {
    return createCookieSessionStorage<ZberatelSession>({ cookie });
  } else if (process.env.SESSION_STORAGE === 'filesystem') {
    return createFileSessionStorage<ZberatelSession>({
      cookie,
      dir: '.sessions',
    });
  }
  return createMemorySessionStorage<ZberatelSession, ZberatelSession>({
    cookie,
  });
};

const { getSession, commitSession, destroySession } =
  getCreateSessiontStorage();

const getSessionFromRequest = async (request: Request) => {
  return getSession(request.headers.get('Cookie'));
};

export async function createUserSession({
  request,
  session,
  remember,
  redirectTo,
}: {
  request: Request;
  session: NhostSession;
  remember: boolean;
  redirectTo: string;
}) {
  const cookie = await getSessionFromRequest(request);
  cookie.set('auth', session);
  return redirect(redirectTo, {
    headers: {
      'Set-Cookie': await commitSession(cookie, {
        maxAge: remember
          ? 60 * 60 * 24 * 7 // 7 days
          : undefined,
      }),
    },
  });
}

export { getSession, commitSession, destroySession, getSessionFromRequest };
