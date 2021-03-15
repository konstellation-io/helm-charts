import logging
from pathlib import Path

import numpy as np
import pandas as pd
import transformers
from transformers import PreTrainedModel, PreTrainedTokenizer

from exceptions import AssetLoadingException


class AssetLoader:
    """
    AssetLoader is responsible of loading and checking for integrity all the assets needed for
    the data pipeline.
    """

    def __init__(self, path: str):
        self.log = logging.getLogger("AssetLoader")
        self.path = path
        self.dataset = self._load_dataset()
        self.model = self._load_model()
        self.vectors = self._load_dataset_vectors()
        self.tokenizer = self._load_tokenizer()
        self._asset_checks()

    def _asset_checks(self):
        """
        Integrity checks for all the assets.
        """
        if not len(self.dataset) == len(self.vectors):
            message = (
                f"The specified dataset (n={len(self.dataset)}) "
                f"and the vectors (available for {len(self.vectors)} documents) do not match. "
                "Please check the inputs."
            )
            raise AssetLoadingException(message)

    def _load_dataset(self) -> pd.DataFrame:
        """
        Loads the dataset from the filepath specified in object attributes.
        For current usage, this must be identical to the training set on which self.model was trained on.
        """
        path = Path(self.path, "dataset.pkl.gz")
        self.log.debug(f"Loading dataset from: {path}")
        df = pd.read_pickle(path, compression="gzip")

        return df

    def _load_dataset_vectors(self) -> np.ndarray:
        """
        Loads vectors from the dataset papers computed using the model.
        """
        path = Path(self.path, "vectors.npy")
        self.log.debug(f"Loading vectors from: {path}")

        return np.load(str(path))

    def _load_model(self) -> PreTrainedModel:
        """
        Loads a Transformer model object from a directory.
        """
        path = Path(self.path, "model")
        self.log.debug(f"Loading model from: {path}")
        model = transformers.AutoModel.from_pretrained(path)

        return model

    def _load_tokenizer(self) -> PreTrainedTokenizer:
        """
        Loads a Transformer tokenizer object from a directory.
        """
        path = Path(self.path, "model")
        self.log.debug(f"Loading tokenizer from: {path}")
        tokenizer = transformers.AutoTokenizer.from_pretrained(path)

        return tokenizer