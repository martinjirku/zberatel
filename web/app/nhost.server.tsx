import { NhostClient } from '@nhost/nhost-js';
import invariant from 'tiny-invariant';

invariant(
  process.env.NEXT_PUBLIC_NHOST_ADMIN_SECRET,
  'NEXT_PUBLIC_NHOST_ADMIN_SECRET must be set'
);

export const createNhostClient = () => {
  return new NhostClient({
    subdomain: process.env.NEXT_PUBLIC_NHOST_BACKEND ?? 'localhost',
    region: process.env.NEXT_PUBLIC_NHOST_REGION,
    adminSecret:
      process.env.NEXT_PUBLIC_NHOST_ADMIN_SECRET ?? 'nhost-admin-secret',
  });
};
