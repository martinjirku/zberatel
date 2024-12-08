import { PropsWithChildren } from "react";
import { twMerge } from "tailwind-merge";

export default function Table({
  children,
  className,
  ...props
}: PropsWithChildren<JSX.IntrinsicElements["table"]>) {
  return (
    <table
      {...props}
      className={twMerge(
        "text-sm border-surface border font-medium text-foreground dark:bg-surface-dark",
        className,
      )}
    >
      {children}
    </table>
  );
}

export const THead = ({
  children,
  className,
  ...props
}: PropsWithChildren<JSX.IntrinsicElements["thead"]>) => {
  return (
    <thead
      {...props}
      className={twMerge(
        "border-b border-surface bg-surface-light text-foreground dark:bg-surface-dark",
        className,
      )}
    >
      {children}
    </thead>
  );
};
export const TBody = ({
  children,
  className,
  ...props
}: PropsWithChildren<JSX.IntrinsicElements["tbody"]>) => {
  return (
    <tbody
      {...props}
      className={twMerge("group text-sm text-black dark:text-white", className)}
    >
      {children}
    </tbody>
  );
};
export const TD = ({
  children,
  className,
  ...props
}: PropsWithChildren<JSX.IntrinsicElements["td"]>) => {
  return (
    <th
      {...props}
      className={twMerge(
        "text-start font-medium border-r last:border-0",
        className,
      )}
    >
      {children}
    </th>
  );
};
export const TH = ({
  children,
  className,
  ...props
}: PropsWithChildren<JSX.IntrinsicElements["th"]>) => {
  return (
    <th
      {...props}
      className={twMerge(
        "text-start font-medium border-r last:border-0",
        className,
      )}
    >
      {children}
    </th>
  );
};
export const TR = ({
  children,
  className,
  ...props
}: PropsWithChildren<JSX.IntrinsicElements["tr"]>) => {
  return (
    <tr
      {...props}
      className={twMerge("border-b border-surface last:border-0", className)}
    >
      {children}
    </tr>
  );
};

Table.tr = TR;
Table.td = TD;
Table.th = TH;
Table.Body = TBody;
Table.Head = THead;
