import { Layer } from "../domain/layer";

export const layerToDisplayName = (layer: Layer) =>
  layer === Layer.Two ? "LayerX L2" : "LayerX L3";
