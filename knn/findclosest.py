def find_closest_k(
    A: np.ndarray, B: np.ndarray, k: int, distance_metric: str = "euclidean"
):
    """
    For each b in B, we try to find the closest k observations in A, given a
    metric.
    """
    if k < 1:
        raise ValueError("k must be positive")

    b_len = B.shape[0]  # number of rows, find k nearest neighbor for each one
    indices = np.zeros((b_len, k), dtype=int)
    distances = np.zeros((b_len, k))

    for i, b in enumerate(B):
        # we need to calculate distances for the whole data set to find the top
        # k nearest ones
        match distance_metric:
            case "euclidean_squared":
                dists = np.sum((A - b) ** 2, axis=1)
            case "euclidean":
                dists = np.sqrt(np.sum((A - b) ** 2, axis=1))
            case "manhattan":
                dists = np.sum(np.abs(A - b), axis=1)
            case _:
                raise ValueError(f"invalid metric: {distance_metric}")

        # argpartition T(n) is more efficient than argsort T(n log n), as it
        # only does a partial sort. As k is likely small compared to the whole
        # dataset, use argpartition with a separate sort on a smaller set
        # instead of argsort on the whole data.
        k_indices = np.argpartition(dists, k - 1)[:k]
        k_distances = dists[k_indices]

        # this is now working on a smaller subset
        sort_order = np.argsort(k_distances)

        # collect the indices and the distances
        indices[i] = k_indices[sort_order]
        distances[i] = k_distances[sort_order]

    return indices, distances
