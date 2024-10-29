#!/usr/bin/env python

"""
Image munging.
"""

import collections
import imageio.v2 as imageio
import skimage
import numpy as np
import random

Dim = collections.namedtuple("Dim", "w h")

files = [
    "tv.png",
    "lg.png",
    "24.png",
    "tr3.png",
]


def autopad(img, dim=Dim(800, 600)):
    """
    Autopad creates an image of a given size and will center the given image
    into it.
    """
    h, w, c = img.shape
    if c == 3:
        img = np.dstack([img, np.zeros((h, w, 1))])
    if w > dim.w or h > dim.h:
        raise ValueError(f"downsizing not yet supported: w={w} h={h}")
    norm_img = np.zeros([dim.h, dim.w, 4], dtype=np.uint8)
    norm_img[:, :, 3] = 255
    r = (dim.h - h) // 2
    c = (dim.w - w) // 2
    norm_img[r : r + h, c : c + w, :] = img
    return norm_img


def make_empty_image(dim=Dim(800, 600)):
    img = np.zeros([dim.h, dim.w, 4], dtype=np.uint8)
    img[:, :, 3] = 255
    return img


def noisy(img):
    alpha = img[:, :, 3]
    rgb = img[:, :, :3]
    noisy = skimage.util.random_noise(rgb, "gaussian", var=random.random())
    noisy = (noisy * 255).astype(np.uint8)
    return np.dstack([noisy, alpha])


def main():
    images = []
    empty_image = make_empty_image()
    for i, fn in enumerate(files):
        img = imageio.imread(fn)
        img = autopad(img)

        if i == len(files) - 1:
            images.extend([noisy(img) for _ in range(7)])
            images.extend([img for _ in range(60)])
        else:
            images.extend([noisy(img) for _ in range(7)])
            images.extend([img for _ in range(5)])
            images.extend([noisy(img) for _ in range(7)])
            images.extend([empty_image for _ in range(1)])

    imageio.mimsave("anim.gif", images, fps=7, loop=0)


if __name__ == "__main__":
    main()
