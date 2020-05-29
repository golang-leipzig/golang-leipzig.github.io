---
title: "Virtual Meetup #11 wrap-up"
date: 2020-05-29T20:00:00+02:00
draft: false
tags:
- summary
- meetup
---

{{% h2 %}}Kubeflow{{% /h2 %}}

The data science process is not free of friction, especially when acquiring,
preparing and cleaning data. But even when data is readily available, one might
want to use it to train a number of ML model to perform model selection.

And what happens, if an update is available and the whole process needs to be
run again. Optimally with a tested model that is put into production in
a canary deployment workflow?

The [Kubeflow](https://www.kubeflow.org/) project is a cloud native machine
learning toolkit supporting experimentation, pipeline building and deployment.
As a cloud native tool it is
[composed](https://www.kubeflow.org/docs/components/) of a set of deployments
and services, from managing Jupyter notebooks, pipelines, dashboard and serving
tools.

[Sascha](github.com/saschagrunert/) and [Markus](https://github.com/mbu93/)
took us on an amazing tour using a real world use case, showcasing the various
parts of Kubeflow and how they fit together.

Rather than repeating the details here, check out the blog post [My exciting
journey into Kubernetes’
history](https://kubernetes.io/blog/2020/05/my-exciting-journey-into-kubernetes-history/):

> A story of data science-ing 90,000 GitHub issues and pull requests by using Kubeflow, TensorFlow, Prow and a fully automated CI/CD pipeline.

The accompanying repository can be found at: [https://github.com/kubernetes-analysis/kubernetes-analysis](https://github.com/kubernetes-analysis/kubernetes-analysis).
