# Important Disclaimers

~This is a rough Google-translated and lightly edited by a human (E. Dolstra). Please verify that the English is fluent and accurate, preserving the quotes exactly and without adding any additional new information.

I will now produce the summary. Note that since the provided text consists of an abstract plus a Dutch summary, I will primarily focus on the English content, as that is the bulk of the thesis.

***

# Summary of *The Purely Functional Software Deployment Model*

## Research Question and Problem Statement

This thesis investigates the problem of correct software deployment in heterogeneous distributed systems. It frames software deployment problem as the inability to guarantee that installing a component (e.g., a library) produced on one machine will function correctly on another. Traditional deployment tools (e.g., RPM, Gentoo, .NET) suffer from issues such as incomplete dependency declarations, interference between versions, lack of rollbacks, and a separation between source and binary deployment.

## Approach

Nix addresses these issues through a **purely functional deployment model**.

**Approach and Method**

Nix provides a **purely functional deployment model**. This model treats component builds as deterministic functions of their inputs: a component's content is defined solely by its build inputs. This differs from destructive models where components can change after installation or share mutable locations (DLL Hell). Instead, every component is stored in its own isolated location in a **Nix store** (e.g., `/nix/store/bwacc7a5c5n3...-...-hello-2.1.1`). The file name contains a cryptographic hash of all inputs involved in the build (sources, compiler, dependencies, etc.). A change in any input causes the hash to change, creating a unique store path.

This hashing approach prevents **interference** because different build variants automatically coexist without interference.

## Results

**Key Results**
1.  **Complete Dependencies**: The hashing scheme prevents undeclared dependencies. Because build environments are isolated (e.g., empty `PATH` or patched dynamic linkers), a build fails deterministically if undeclared dependencies are present.
2.  **Side-by-Side Deployment**: Different versions or variants (e.g., Firefox 1.0 vs 2.0) coexist.
3.  **Atomic Upgrades and Rollbacks**: Upgrades are **atomic**, with generation-based user environments ensuring users never see an inconsistent state. Old configurations persist, allowing instant rollback via symlink switching.
4.  **Transparent Source/Binary Deployment**: Nix supports a transparent model where binary deployment is an automatic optimization of source deployment. Pre-built components can be registered as substitutes. If available, they are installed; otherwise, Nix falls back to building from source.
5.  **Garbage Collection**: Unused components can be safely removed using garbage collection, which deletes only those not reachable from configuration roots (like active user environments).

## Why it Matters

Nix allows **side-by-side deployment** without interference, ensuring correct behavior regardless of the order upgrades or changes elsewhere in the system. It unifies build and deployment into a **single formalism**, making configuration traceable and reproducible.

However, efficient upgrading remains a primary challenge. While binary patching can reduce network bandwidth, changes to fundamental dependencies (e.g., Glibc) require dependent components to be rebuilt, increasing disk space usage. A full "destructive" upgrade might be more space-efficient but sacrifices reproducibility and rollbacks.

> "The purely functional deployment model implemented in Nix and the cryptographic hashing scheme of the Nix store in particular give us important features that are lacking in most deployment systems, such as complete dependencies, complete deployment, side-by-side deployment, atomic upgrades and rollbacks, transparent source/binary deployment and reproducibility (see Section 1.5)."

> "The cryptographic hashing scheme is effective in preventing undeclared dependencies, assuming a fully Nixified environment."

> "However, there are also a number of reasonable criticisms that can be leveled at the purely functional model... Efficient upgrading remains a problem. Using patch deployment, upgrades can be done efficiently in terms of network bandwidth. But a change to a fundamental dependency can still cause disk space requirements to double. This is a real problem compared to destructive updates. However, it can be argued that disk space is abundant, and that software components no longer dominate disk space consumption."

> "Preventing undeclared dependencies is good, but the lack of scoped composition mechanisms lessens its usefulness (as discussed on page 173)."

> "Efficient upgrading remains a problem. Using patch deployment, upgrades can be done efficiently in terms of network bandwidth. But a change to a fundamental dependency can still cause disk space requirements to double. This is a real problem compared to destructive updates. However, it can be argued that disk space is abundant, and that software components no longer dominate disk space consumption."

> "Preventing undeclared dependencies is good, but the lack of scoped composition mechanisms lessens its usefulness (as discussed on page 173)."

> "Most of our experience with Nix has been on Linux, which allows a self-contained, pure Nixpkgs (page 169). That is, most deployed components can be built and deployed through Nix. This is an ideal situation that cannot be achieved in most other operating systems. A GUI application on Mac OS X will typically have a dependency on the Cocoa or Carbon GUI libraries. Since these are closed source and cannot be copied legally, we can neither build them nor use the techniques of Section 7.1.4 to deploy them in binary form. Thus, we are forced to resort to impurity (e.g., by linking against the non-store path /System/Library/Frameworks/Cocoa.framework , which reduces Nix's effectiveness."

> "Nix combines the best of both worlds... In essence, binary deployment can be considered a partial evaluation of source deployment with respect to a specific platform, and such optimisations should be 'invisible'."

> "This is a serious problem that needs to be solved. If a fundamental component (say, Zlib) contains a security bug [1]. This component is used by many other components. Thus, we want to ensure that we upgrade all installed applications in all user environments to new instances that do not use the bad component anymore. Of course, we can just run nix-env -u "*" ... and hope that this gets rid of all bad components. But there is no guarantee that this will succeed... Instead, the purely functional model's non-interference property now bites us. Thus, secure sharing is not possible between users who use the same remote trust relations. Thus, there should be a generic mechanism that detects uses of bad components... Since there is no guarantee that any particular substitution will succeed if all other users trust source X . . Nor can there be such thing as safe, since this is not a problem since that is why we have such things as fixed versions."

> "Since the purely functional model's non-interference property now bites us. Thus, secure sharing is not possible between users who use the same remote trust relations. This is specifically the case in the extensional model, where we cannot allow users to have write
