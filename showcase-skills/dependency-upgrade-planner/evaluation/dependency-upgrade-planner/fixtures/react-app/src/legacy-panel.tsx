import React from "react";

export function LegacyPanel({ title }: { title?: string }) {
  const element = React.createFactory("section");
  return element(null, title ?? "Overview");
}

LegacyPanel.defaultProps = { title: "Overview" };
